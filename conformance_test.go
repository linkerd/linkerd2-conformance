package conformance_test

import (
	"testing"

	"github.com/linkerd/linkerd2-conformance/specs/install"
	"github.com/linkerd/linkerd2-conformance/specs/uninstall"
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	utils.InitTestHelper()
	_ = m.Run()
}

func TestConformance(t *testing.T) {

	// A Describe block to hold the tests
	_ = ginkgo.Describe("", func() {
		// Bring tests into scope
		_ = install.NewInstallSpec()
		_ = uninstall.NewUninstallSpec()
	})

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "linkerd2 conformance validation")
}
