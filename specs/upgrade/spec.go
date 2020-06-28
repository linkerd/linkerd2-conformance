package upgrade

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// RunUpgradeSpec runs upgrade tests
func RunUpgradeSpec() bool {
	return ginkgo.Describe("`linkerd upgrade`", func() {
		h, _ := utils.GetHelperAndConfig()

		if h.UpgradeFromVersion() == "" {
			ginkgo.Skip("Skipping upgrade tests")
		}

		ginkgo.It("can upgrade CLI", testUpgradeCLI)
		ginkgo.It("can upgrade control-plane", testUpgrade)

		// TODO: add tests for upgrading from manifests
		// TODO: add test for upgrading using 2 stage installation
	})
}
