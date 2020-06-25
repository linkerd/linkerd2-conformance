package uninstall

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// RunUninstallSpec runs the uninstall test suite
func RunUninstallSpec() bool {
	return ginkgo.Describe("`linkrd install`", func() {
		ginkgo.It("can uninstall control plane", func() {
			h := utils.TestHelper
			utils.UninstallLinkerdControlPlane(h)
		})
	})
}
