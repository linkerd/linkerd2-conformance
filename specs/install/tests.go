package install

import (
	"fmt"

	"github.com/linkerd/linkerd2/testutil"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func testControlPlaneInstall(h *testutil.TestHelper) {

	if err := h.CheckIfNamespaceExists(h.GetLinkerdNamespace()); err == nil {
		ginkgo.Skip(fmt.Sprintf("linkerd control plane already exists in namespace %s", h.GetLinkerdNamespace()))
	}

	ginkgo.By("verifying if Helm release is empty")
	gomega.Expect(h.GetHelmReleaseName()).To(gomega.Equal(""))

	cmd := "install"
	args := []string{
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
}

func testControlPlaneState(h *testutil.TestHelper) {
	// Test Namespace
	linkerdNs := h.GetLinkerdNamespace()
	ginkgo.By(fmt.Sprintf("checking if %s namespace exists", linkerdNs))
	err := h.CheckIfNamespaceExists(linkerdNs)
	gomega.Expect(err).To(gomega.BeNil())

	// TestServices
	linkerdSvcs := []string{
		"linkerd-controller-api",
		"linkerd-dst",
		"linkerd-grafana",
		"linkerd-identity",
		"linkerd-prometheus",
		"linkerd-web",
		"linkerd-tap",
	}

	for _, svc := range linkerdSvcs {
		err = h.CheckService(linkerdNs, svc)
		gomega.Expect(err).To(gomega.BeNil())
	}

	// Test Pods and Deployments
	for deploy, spec := range testutil.LinkerdDeployReplicas {
		err = h.CheckPods(linkerdNs, deploy, spec.Replicas)
		gomega.Expect(err).To(gomega.BeNil())

		err = h.CheckDeployment(linkerdNs, deploy, spec.Replicas)
		gomega.Expect(err).To(gomega.BeNil())
	}
}
