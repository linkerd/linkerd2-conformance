package install

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// RunInstallSpec runs the install test suite
func RunInstallSpec() bool {
	return ginkgo.Describe("`linkerd install`", func() {
		h := utils.TestHelper
		ha := utils.TestConfig.HA()

		ginkgo.It("can install a new control plane", func() {
			utils.InstallLinkerdControlPlane(h, ha)
		})
	})
}
