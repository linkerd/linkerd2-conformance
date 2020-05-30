package install

import (
	"fmt"

	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/linkerd/linkerd2/testutil"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Verifying linkerd2 control plane components", func() {
	var cmd string
	var args []string
	var h *testutil.TestHelper

	ginkgo.BeforeEach(func() {
		h = utils.TestHelper
	})

	ginkgo.Context("By attempting to installing control plane components", func() {
		ginkgo.It("should successfully pipe output of `linkerd install` to `kubectl`", func() {

			if err := h.CheckIfNamespaceExists(h.GetLinkerdNamespace()); err == nil {
				ginkgo.Skip(fmt.Sprintf("linkerd control plane already exists in namespace %s", h.GetLinkerdNamespace()))
			}

			ginkgo.By("Verifying if Helm release is empty")
			gomega.Expect(h.GetHelmReleaseName()).To(gomega.Equal(""))

			cmd = "install"
			args = []string{
				"--controller-log-level", "debug",
				"--proxy-log-level", "warn,linkerd2_proxy=debug",
				"--proxy-version", h.GetVersion(),
			}
			if h.GetClusterDomain() != "cluster.local" {
				args = append(args, "--cluster-domain", h.GetClusterDomain())
			}

			// TODO: handle external issuer

			// TODO: handle upgrade

			exec := append([]string{cmd}, args...)

			ginkgo.By("attempting to issue `linkerd install` and gather the manifests")
			out, stderr, _ := h.LinkerdRun(exec...)
			gomega.Expect(stderr).To(gomega.Equal(""))

			ginkgo.By("attempting to apply manifests to your cluster")
			out, err := h.KubectlApply(out, "")

			gomega.Expect(err).To(gomega.BeNil())

		})

	})
})
