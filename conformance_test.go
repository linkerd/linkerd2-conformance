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

	h, c := utils.GetHelperAndConfig()

	// install linkerd binary

	version := h.UpgradeFromVersion()
	if version == "" {
		version = h.GetVersion()
	}

	if err := utils.InstallLinkerdBinary(c.GetLinkerdPath(), version, false, true); err != nil {
		fmt.Printf("error installing linkerd2 (%s): %s", version, err.Error())
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func TestConformance(t *testing.T) {
	specs.RunAllSpecs(t)
}
