package inject

import (
	"fmt"

	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// RunInjectSpec runs inject tests
func RunInjectSpec() bool {
	return ginkgo.Describe("`linkerd inject`", func() {
		if skip := utils.TestConfig.Inject.SkipTest; skip {
			ginkgo.Skip(fmt.Sprintf("Skipping inject tests: inject.skil set to \"%v\" in config YAML", skip))
		}
		ginkgo.It("can perform manual injection", func() {

			ginkgo.When("without parameters", func() {
				testInjectManual(false)
			})

			ginkgo.Context("with parameters", func() {
				testInjectManual(true)
			})
		})

		ginkgo.It("can inject proxy container into pods", testProxyInjection)

		if clean := utils.TestConfig.Inject.Clean; clean {
			ginkgo.It("should delete all resources created during testing", testClean)
		}
	})
}
