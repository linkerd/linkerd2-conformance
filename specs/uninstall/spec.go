package uninstall

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// RunUninstallSpec runs the tests for `linkerd uninstall`
func RunUninstallSpec() bool {
	return ginkgo.Describe("`linkerd uninstall`", func() {
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
