package check

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

func RunCheckSpec(pre bool) bool {
	return ginkgo.Describe("`linkerd check`", func() {
		h, _ := utils.GetHelperAndConfig()
		if pre {
			ginkgo.It("can successfully run all pre-installation checks", func() {
				utils.RunCheck(h, true)
			})
		} else {
			ginkgo.It("can successfully run all post-installation checks", func() {
				utils.RunCheck(h, false)
			})
		}
	})
}
