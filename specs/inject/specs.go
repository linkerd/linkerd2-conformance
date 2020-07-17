package inject

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// RunInjectTests runs the specs for inject
func RunInjectTests() bool {
	return ginkgo.Describe("`linkerd inject`", func() {
		_, c := utils.GetHelperAndConfig()

		_ = utils.ShouldTestSkip(c.SkipInject(), "Skipping inject tests")

		ginkgo.It("can perform manual injection", func() {

			ginkgo.When("without parameters", func() {
				testInjectManual(false)
			})

			ginkgo.When("with parameters", func() {
				testInjectManual(true)
			})
		})

		ginkgo.It("can inject proxy container into pods", testProxyInjection)

		ginkgo.It("can override pod level proxy config with namespace level config", testInjectAutoNsOverrideAnnotations)

		if clean := c.CleanInject(); clean {
			ginkgo.It("should delete all resources created during testing", testClean)
		}
	})
}
