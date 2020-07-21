package ingress

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// RunIngressTests runs the specs for ingress
func RunIngressTests() bool {
	return ginkgo.Describe("ingress: ", func() {
		_, c := utils.GetHelperAndConfig()

		_ = utils.ShouldTestSkip(c.SkipIngress(), "Skipping ingress tests")

		ginkgo.It("can install and inject emojivoto app", func() {
			utils.TestEmojivotoApp()
			utils.TestEmojivotoInject()
		})

		if c.ShouldTestIngressOfType(utils.Nginx) {
			ginkgo.It("can work with nginx ingress controller", testNginx)
		}

		ginkgo.It("can uninstall emojivoto app", utils.TestEmojivotoUninstall)
	})
}
