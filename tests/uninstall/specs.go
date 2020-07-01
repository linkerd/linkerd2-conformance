package uninstall

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

var _ = ginkgo.Describe("`linkerd uninstall`", func() {
	h, c := utils.GetHelperAndConfig()

	ginkgo.BeforeEach(func() {
		if !(c.SingleControlPlane() && h.Uninstall()) {
			ginkgo.Skip("Skipping global uninstall test")
		}
	})

	ginkgo.It("can uninstall the control plane", func() {
		utils.UninstallLinkerdControlPlane(h)
	})
})
