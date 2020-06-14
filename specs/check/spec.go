package check

import (
	"fmt"

	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// RunCheckSpec runs the tests for `check`
func RunCheckSpec(pre bool) bool {
	return ginkgo.Describe("`linkerd check`", func() {
		ginkgo.Context(fmt.Sprintf("with --pre: %v", pre), func() {
			h := utils.TestHelper
			ginkgo.It("should successfully pass all checks", func() {
				testCheck(h, pre)
			})
		})
	})
}
