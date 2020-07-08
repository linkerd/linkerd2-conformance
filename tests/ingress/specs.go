package ingress

import (
	"github.com/onsi/ginkgo"
)

var _ = ginkgo.Describe("linkerd", func() {
	// _, _ := utils.GetHelperAndConfig()

	ginkgo.It("can work with nginx ingress controller", testNginx)
	//	ginkgo.Describe("linkerd", func() {
	//		if c.ShouldTestIngressOfType(utils.Nginx) {
	//			ginkgo.It(fmt.Sprintf("can work with %s ingress controller", utils.Nginx), func() {})
	//		}
	//
	//		if c.ShouldTestIngressOfType(utils.Traefik) {
	//			ginkgo.It(fmt.Sprintf("can work with %s ingress controller", utils.Traefik), func() {})
	//		}
	//
	//		if c.ShouldTestIngressOfType(utils.Ambassador) {
	//			ginkgo.It(fmt.Sprintf("can work with %s ingress controller", utils.Ambassador), func() {})
	//		}
	//
	//		if c.ShouldTestIngressOfType(utils.Gloo) {
	//			ginkgo.It(fmt.Sprintf("can work with %s ingress controller", utils.Gloo), func() {})
	//		}
	//
	//		if c.ShouldTestIngressOfType(utils.GCE) {
	//			ginkgo.It(fmt.Sprintf("can work with %s ingress controller", utils.GCE), func() {})
	//		}
	//
	//		if c.ShouldTestIngressOfType(utils.Contour) {
	//			ginkgo.It(fmt.Sprintf("can work with %s ingress controller", utils.Contour), func() {})
	//
	//		}
	//	})
})
