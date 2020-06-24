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

var (
	// TestHelper stores a global instace of tesutil.TestHelper
	TestHelper *testutil.TestHelper

	// TestConfig stores a global instance of the parsed config YAML
	TestConfig *ConformanceTestOptions
)

const (
	installEnv           = "LINKERD2_VERSION"
	configFile           = "config.yaml"
	linkerdInstallScript = "install.sh"
	installScriptURL     = "https://raw.githubusercontent.com/linkerd/website/master/run.linkerd.io/public/install"
)

// InitTestHelper initializes a test helper
func InitTestHelper() error {

	var err error

	if fileExists(configFile) {
		yamlFile, err := ioutil.ReadFile(configFile)
		if err != nil {
			return err
		}

		if err := yaml.Unmarshal(yamlFile, &TestConfig); err != nil {
			return err
		}

	}

	if err := TestConfig.parseDefaultConfigValues(); err != nil {
		return err
	}

	TestHelper, err = TestConfig.initNewTestHelperFromOptions()
	if err != nil {
		return err
	}

	if err = installLinkerdIfNotExists(TestConfig.LinkerdBinaryPath, TestConfig.LinkerdVersion); err != nil {
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

// Err returns err.Error() string
// if err is not nil
// This helper is meant to be used with
// gomega.Should() to annotate failures
// without causing runtime errors when
// returned errors are nil
func Err(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
