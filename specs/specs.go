package specs

import (
	"testing"

	"github.com/linkerd/linkerd2-conformance/specs/check"
	"github.com/linkerd/linkerd2-conformance/specs/install"
	"github.com/linkerd/linkerd2-conformance/specs/uninstall"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func RunAllSpecs(t *testing.T) {

	// A single top-level wrapper Describe is required to prevent
	// the specs from being run in a random order
	// The Describe message is intentionally left empty
	// as it only serves to prevent randomization of specs
	_ = ginkgo.Describe("", func() {

		// Bring tests into scope
		_ = check.RunCheckSpec(true)     // pre checks
		_ = install.RunInstallSpec()     // install
		_ = check.RunCheckSpec(false)    // post checks
		_ = uninstall.RunUninstallSpec() // uninstall

		// TODO: The install/uninstall/check specs may have to be moved to
		// `BeforeSuite` depending on how we add the main tests

		// TODO: The order of these checks is not final. As we start adding the more important checks,
		// we may want to run the install, uninstall and check tests for each of the tests.
		// Further, the test config file may also have an option to enable global install / uninstall.
	})

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "linkerd2 conformance validation")
}
