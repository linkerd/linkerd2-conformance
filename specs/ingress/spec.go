package ingress

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

// RunIngressTests runs the specs for ingress
func RunIngressTests() bool {
	return ginkgo.Describe("linkerd", func() {
		_, c := utils.GetHelperAndConfig()

		_ = utils.ShouldTestSkip(c.SkipIngress(), "Skipping ingress tests")

		if c.ShouldTestIngressOfType(utils.Nginx) {
			ginkgo.It("can work with nginx ingress controller", testNginx)
		}
	})
}
