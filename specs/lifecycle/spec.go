package lifecycle

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// RunLifecycleTest runs the specs for lifecycle tests
func RunLifecycleTest() bool {
	return ginkgo.Describe("lifecycle: ", func() {
		h, c := utils.GetHelperAndConfig()

		_ = utils.ShouldTestSkip(c.SkipLifecycle(), "Skipping lifecycle tests")

		ginkgo.Describe("`linkerd install`", func() {
			ginkgo.It("can install a new control plane", func() {
				utils.InstallLinkerdControlPlane(h, c)
			})
		})

		if h.UpgradeFromVersion() != "" {
			ginkgo.Describe("`linkerd upgrade`", func() {
				ginkgo.It("can upgrade CLI", testUpgradeCLI)
				ginkgo.It("can upgrade control-plane", testUpgrade)
			})
		}

		// If each test will have its own control plane, uninstall the currently
		// control plane right away. Else, wait for all tests to complete, and
		// call the uninstall test separately at the end
		if !c.SingleControlPlane() {
			_ = RunUninstallTest()
		}

	})
}

// RunUninstallTest runs the uninstall test separately
func RunUninstallTest() bool {
	return ginkgo.Describe("`linkerd install`", func() {
		h, _ := utils.GetHelperAndConfig()

		ginkgo.It("can uninstall control plane", func() {
			utils.UninstallLinkerdControlPlane(h)
		})
	})
}
