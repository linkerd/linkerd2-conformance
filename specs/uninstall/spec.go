package uninstall

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// NewUninstallSpec returns a new spec for linkerd uninstall
func NewUninstallSpec() bool {
	return ginkgo.Describe("uninstall process", func() {
		h := utils.TestHelper

		ginkgo.BeforeEach(func() {
			if !h.Uninstall() {
				ginkgo.Skip("Skipping uninstall")
			}
		})

		ginkgo.It("should remove the installed control plane", func() {
			testControlPlaneUninstall(h)
		})

	})
}
