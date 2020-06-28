package specs

import (
	"testing"

	"github.com/linkerd/linkerd2-conformance/specs/inject"
	"github.com/linkerd/linkerd2-conformance/specs/install"
	"github.com/linkerd/linkerd2-conformance/specs/uninstall"
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/linkerd/linkerd2/testutil"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

// Run Install / Uninstall test in a separate
// Describe block before running the primary tests
// This is done so that the BeforeEach and AfterEach
// blocks do not interfere with these tests
func runPreFlightSpecs(h *testutil.TestHelper, c *utils.ConformanceTestOptions) bool {
	return ginkgo.Describe("", func() {
		if !c.GlobalControlPlane() && c.SkipInstall() {
			ginkgo.Skip("Skipping `linkerd install` spec")
		}
		_ = install.RunInstallSpec()
		if !c.GlobalControlPlane() { // Immediately uninstall if each test shall have its own control-plane
			_ = uninstall.RunUninstallSpec()
		}
	})
}

func runMainSpecs(h *testutil.TestHelper, c *utils.ConformanceTestOptions) bool {
	return ginkgo.Describe("", func() {
		if !c.GlobalControlPlane() {
			_ = ginkgo.BeforeEach(func() {
				utils.InstallLinkerdControlPlane(h, c.HA())
			})

			_ = ginkgo.AfterEach(func() {
				utils.UninstallLinkerdControlPlane(h)
			})

		}

		// Bring main tests into scope
		_ = inject.RunInjectSpec()

		// global uninstall (if true) should always run at the end
		if c.GlobalControlPlane() && h.Uninstall() {
			_ = uninstall.RunUninstallSpec()
		}
	})
}

// RunAllSpecs wraps all the specs into a single runnable test
func RunAllSpecs(t *testing.T) {

	h, c := utils.GetHelperAndConfig()

	// A single top-level wrapper Describe is required to prevent
	// the specs from being run in a random order
	// The Describe message is intentionally left empty
	// as it only serves to prevent randomization of specs
	_ = ginkgo.Describe("", func() {
		_ = runPreFlightSpecs(h, c)
		_ = runMainSpecs(h, c)
	})

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "linkerd2 conformance validation")
}
