package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/linkerd/linkerd2/testutil"
	"gopkg.in/yaml.v2"
)

var TestHelper *testutil.TestHelper

const install_env = "LINKERD2_VERSION"

func InitTestHelper() error {

	var opt *ConformanceTestOptions
	var err error

	if fileExists("config.yaml") {
		yamlFile, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			return err
		}

		if err := yaml.Unmarshal(yamlFile, &opt); err != nil {
			return err
		}

		if err := opt.parseConfigValues(); err != nil {
			return err
		}

	} else {
		opt, err = getDefaultConformanceOptions()
		if err != nil {
			return err
		}
	}

	TestHelper = opt.initNewTestHelperFromOptions()
	err = initK8sHelper(opt, TestHelper)
	if err != nil {
		return err
	}

	err = installLinkerdIfNotExists(opt.LinkerdBinaryPath, opt.LinkerdVersion)
	if err != nil {
		return err
	}

	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func initK8sHelper(options *ConformanceTestOptions, h *testutil.TestHelper) error {
	k8sHelper, err := testutil.NewKubernetesHelper(options.K8sContext, h.RetryFor)
	if err != nil {
		return fmt.Errorf("Error creating K8s helper: %s", err)
	}
	h.KubernetesHelper = *k8sHelper

	return nil
}

func fetchInstallScript() ([]byte, error) {

	url := "https://raw.githubusercontent.com/linkerd/website/master/run.linkerd.io/public/install"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func makeScriptFile(script []byte, path string) error {
	file, err := os.Create(path)
	defer file.Close()

	_, err = file.Write(script)
	if err != nil {
		return err
	}
	return nil
}

func installLinkerdIfNotExists(linkerd, version string) error {
	if fileExists(linkerd) {
		fmt.Printf("linkerd2 binary exists in \"%s\"- skipping installation\n", linkerd)
		return nil
	}

	file := "install.sh"

	script, err := fetchInstallScript()
	if err != nil {
		return err
	}

	err = makeScriptFile(script, file)
	if err != nil {
		return err
	}

	os.Setenv(install_env, version)

	cmd := exec.Command("/bin/sh", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	err = os.Remove(file)
	if err != nil {
		return err
	}

	return nil
}
