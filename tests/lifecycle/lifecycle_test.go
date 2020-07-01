package lifecycle

import (
	"fmt"
	"os"
	"testing"

	"github.com/linkerd/linkerd2-conformance/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
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

func TestLifecycle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lifecycle")
}
