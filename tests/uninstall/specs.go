package uninstall

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

var _ = ginkgo.Describe("`linkerd uninstall`", func() {
	h, c := utils.GetHelperAndConfig()

	_ = utils.ShouldTestSkip(!(c.SingleControlPlane() && h.Uninstall()), "Skipping uninstall test")

	ginkgo.It("can uninstall the control plane", func() {
		utils.UninstallLinkerdControlPlane(h)
	})
})
