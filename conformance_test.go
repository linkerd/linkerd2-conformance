package conformance_test

import (
	"testing"

	"github.com/linkerd/linkerd2-conformance/specs/check"
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

		// TODO: These tests may have to be moved to `BeforeSuite` depending on how we add the upcoming tests

		_ = check.NewCheckSpec(true)  // pre checks
		_ = install.NewInstallSpec()  // install
		_ = check.NewCheckSpec(false) // post checks

		_ = uninstall.NewUninstallSpec() // uninstall

		// TODO: The order of these checks is not final. As we start adding the more important checks,
		// we may want to run the install, uninstall and check tests for each of the tests.
		// Further, the test config file may also have an option to enable global install / uninstall.
	})

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "linkerd2 conformance validation")
}
