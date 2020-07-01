package lifecycle

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

var _ = ginkgo.Describe("", func() {
	h, c := utils.GetHelperAndConfig()

	ginkgo.BeforeEach(func() {
		if c.SkipLifecycle() {
			ginkgo.Skip("Skipping lifecycle tests")
		}
	})

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

	if !c.SingleControlPlane() {
		ginkgo.Describe("`linkerd install`", func() {
			ginkgo.It("can uninstall control plane", func() {
				utils.UninstallLinkerdControlPlane(h)
			})
		})
	}

})
