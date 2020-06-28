package conformance_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/linkerd/linkerd2-conformance/specs"
	"github.com/linkerd/linkerd2-conformance/utils"
)

func TestMain(m *testing.M) {
	err := utils.InitTestHelper()
	if err != nil {
		fmt.Printf("error initializing tests: %s\n", err.Error())
		os.Exit(1)

	}

	h := utils.TestHelper
	if h.UpgradeFromVersion() == "" { // directly install Linkerd binary if upgrade tests are not going to run
		if err := utils.InstallLinkerdBinary(utils.TestConfig.LinkerdBinaryPath, h.GetVersion(), false); err != nil {
			fmt.Printf("error installing linkerd2 (%s): %s", h.GetVersion(), err.Error())
			os.Exit(1)
		}

	}
	code := m.Run()
	os.Exit(code)
}

func TestConformance(t *testing.T) {
	specs.RunAllSpecs(t)
}
