package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/linkerd/linkerd2/testutil"
	"gopkg.in/yaml.v2"
)

var TestHelper *testutil.TestHelper

func InitTestHelper() error {
	var opt *ConformanceTestOptions

	if fileExists("config.yaml") {
		yamlFile, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(yamlFile, &opt)
		if err != nil {
			return err
		}
		opt.parseConfigValues()

	} else {
		opt = getDefaultConformanceOptions()
	}

	TestHelper = opt.initNewTestHelperFromOptions()
	err := initK8sHelper(opt, TestHelper)
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
