package check

import (
	"fmt"

	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// NewCheckSpec returns a new check test spec
func NewCheckSpec(pre bool) bool {
	return ginkgo.Context("`linkerd check`", func() {
		ginkgo.Context(fmt.Sprintf("With --pre: %v", pre), func() {
			h := utils.TestHelper
			ginkgo.It("should successfully pass all checks", func() {
				testCheck(h, pre)
			})
		})
	})
}
