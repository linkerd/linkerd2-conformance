package uninstall

import (
	"github.com/linkerd/linkerd2/testutil"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func testControlPlaneUninstall(h *testutil.TestHelper) {

	cmd := "install"
	args := []string{
		"--ignore-cluster",
	}

	exec := append([]string{cmd}, args...)

	ginkgo.By("gathering control plane manigfests and piping to `kubectl delete`")
	out, stderr, err := h.LinkerdRun(exec...)
	gomega.Expect(stderr).To(gomega.Equal(""))

	args = []string{"delete", "-f", "-"}

	ginkgo.By("attempting to delete resources to your cluster")
	out, err = h.Kubectl(out, args...)
	gomega.Expect(err).To(gomega.BeNil())
}
