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

// TestHelper stores an instace of tesutil.TestHelper
var TestHelper *testutil.TestHelper

const (
	installEnv           = "LINKERD2_VERSION"
	configFile           = "config.yaml"
	linkerdInstallScript = "install.sh"
	installScriptURL     = "https://raw.githubusercontent.com/linkerd/website/master/run.linkerd.io/public/install"
)

// InitTestHelper initializes a test helper
func InitTestHelper() error {

	var opt *ConformanceTestOptions
	var err error

	if fileExists(configFile) {
		yamlFile, err := ioutil.ReadFile(configFile)
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

	TestHelper, err = opt.initNewTestHelperFromOptions()
	if err != nil {
		return err
	}

	if err = installLinkerdIfNotExists(opt.LinkerdBinaryPath, opt.LinkerdVersion); err != nil {
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

func fetchInstallScript() ([]byte, error) {

	req, err := http.NewRequest("GET", installScriptURL, nil)
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

	if _, err = file.Write(script); err != nil {
		return err
	}
	return nil
}

func installLinkerdIfNotExists(linkerd, version string) error {
	if fileExists(linkerd) {
		fmt.Printf("linkerd2 binary exists in \"%s\"- skipping installation\n", linkerd)
		return nil
	}

	script, err := fetchInstallScript()
	if err != nil {
		return err
	}

	if err = makeScriptFile(script, linkerdInstallScript); err != nil {
		return err
	}

	os.Setenv(installEnv, version)

	cmd := exec.Command("/bin/sh", linkerdInstallScript)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	if err = os.Remove(linkerdInstallScript); err != nil {
		return err
	}

	return nil
}
