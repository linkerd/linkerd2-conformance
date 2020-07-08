package ingress

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
)

var _ = ginkgo.Describe("linkerd", func() {
	_, c := utils.GetHelperAndConfig()

	ginkgo.BeforeEach(func() {
		if c.SkipIngress() {
			ginkgo.Skip("Skipping ingress tests")
		}
	})

	if c.ShouldTestIngressOfType(utils.Nginx) {
		ginkgo.It("can work with nginx ingress controller", testNginx)
	}
})
