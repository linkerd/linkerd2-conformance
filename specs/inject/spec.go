package inject

import (
	"github.com/onsi/ginkgo"
)

// RunInjectSpec runs inject tests
func RunInjectSpec() bool {
	return ginkgo.Describe("`linkerd inject`", func() {

		ginkgo.It("can perform manual injection", func() {

			ginkgo.When("without parameters", func() {
				testInjectManual(false)
			})

			ginkgo.Context("with parameters", func() {
				testInjectManual(true)
			})
		})

	})
}
