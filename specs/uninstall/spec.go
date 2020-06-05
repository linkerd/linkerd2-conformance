package uninstall

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// NewUninstallSpec returns a new spec for linkerd uninstall
func NewUninstallSpec() bool {
	return ginkgo.Context("uninstall process", func() {
		h := utils.TestHelper
		ginkgo.It("should remove the installed control plane", func() {
			testControlPlaneUninstall(h)
		})

	})
}
