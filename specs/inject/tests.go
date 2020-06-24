package inject

import (
	"fmt"

	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/linkerd/linkerd2/pkg/k8s"
	"github.com/linkerd/linkerd2/testutil"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var (
	proxyInjectTestNs string
)

func testInjectManual(withParams bool) {
	var golden string

	testHelper := utils.TestHelper
	injectYAMLPath := "testdata/inject/inject_test.yaml"
	cmd := []string{"inject",
		"--manual",
		"--linkerd-namespace=fake-ns",
		"--disable-identity",
		"--ignore-cluster",
		"--proxy-version=proxy-version",
		"--proxy-image=proxy-image",
		"--init-image=init-image",
	}

	if withParams {
		ginkgo.By("Adding manual parameters to `linkerd inject`")
		params := []string{
			"--disable-tap",
			"--image-pull-policy=Never",
			"--control-port=123",
			"--skip-inbound-ports=234,345",
			"--skip-outbound-ports=456,567",
			"--inbound-port=678",
			"--admin-port=789",
			"--outbound-port=890",
			"--proxy-cpu-request=10m",
			"--proxy-memory-request=10Mi",
			"--proxy-cpu-limit=20m",
			"--proxy-memory-limit=20Mi",
			"--proxy-uid=1337",
			"--proxy-log-level=warn",
			"--enable-external-profiles",
		}
		for _, param := range params {
			cmd = append(cmd, param)
		}
		golden = "inject/inject_params.golden"
	} else {
		golden = "inject/inject_default.golden"
	}
	cmd = append(cmd, injectYAMLPath)

	ginkgo.By(fmt.Sprintf("Running `linkerd inject` against %s", injectYAMLPath))
	out, stderr, err := testHelper.LinkerdRun(cmd...)

	gomega.Expect(err).Should(gomega.BeNil(), stderr)

	ginkgo.By("Validating injected output")
	err = testutil.ValidateInject(out, golden, testHelper)
	gomega.Expect(err).To(gomega.BeNil())
}

func testProxyInjection() {
	h := utils.TestHelper

	ginkgo.By("Reading pod YAML")
	podYAML, err := testutil.ReadFile("testdata/inject/pod.yaml")

	gomega.Expect(err).Should(gomega.BeNil(), utils.Err(err))

	injectNs := "inject-pod-test"
	podName := "inject-pod-test-terminus"
	nsAnnotations := map[string]string{
		k8s.ProxyInjectAnnotation: k8s.ProxyInjectEnabled,
	}

	proxyInjectTestNs = h.GetTestNamespace(injectNs)
	ginkgo.By(fmt.Sprintf("Creating data plane namespace %s", proxyInjectTestNs))
	err = h.CreateDataPlaneNamespaceIfNotExists(proxyInjectTestNs, nsAnnotations)

	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to create namespace %s: %s", proxyInjectTestNs, utils.Err(err)))

	ginkgo.By(fmt.Sprintf("Creating test pod in namespace %s", proxyInjectTestNs))
	o, err := h.Kubectl(podYAML, "-n", proxyInjectTestNs, "create", "-f", "-")

	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to create pod/%s in namespace %s for %s: %s", podName, proxyInjectTestNs, utils.Err(err), o))

	ginkgo.By("Waiting for pod to be initialized")
	o, err = h.Kubectl("", "-n", proxyInjectTestNs, "wait", "--for=condition=initialized", "--timeout=120s", "pod/"+podName)

	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to wait for pod/%s to be initialized in namespace %s: %s: %s", podName, proxyInjectTestNs, utils.Err(err), o))

	ginkgo.By(fmt.Sprintf("Getting pods from namespace %s", proxyInjectTestNs))
	pods, err := h.GetPods(proxyInjectTestNs, map[string]string{"app": podName})

	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to get pods in namespace %s", proxyInjectTestNs))
	gomega.Expect(len(pods)).Should(gomega.Equal(1), fmt.Sprintf("found %d pods, expected %d", len(pods), 1))

	containers := pods[0].Spec.Containers
	proxyContainers := testutil.GetProxyContainer(containers)
	gomega.Expect(proxyContainers).ShouldNot(gomega.BeNil(), fmt.Sprint("proxy container not injected"))
}

func testClean() {
	h := utils.TestHelper

	namespaces := []string{
		proxyInjectTestNs,
	}

	for _, ns := range namespaces {
		ginkgo.By(fmt.Sprintf("Gathering manifests for namespace/%s", ns))
		out, err := h.Kubectl("", "-n", ns, "get", "all", "-o", "yaml")

		gomega.Expect(err).Should(gomega.BeNil(), utils.Err(err))

		ginkgo.By(fmt.Sprintf("Deleting resources in namespace/%s", ns))
		_, err = h.Kubectl(out, "delete", "-f", "-")

		gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("could not delete resources in namespace/%s: ", ns, utils.Err(err)))

		_, err = h.Kubectl("", "delete", "ns", ns)

		gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("could not delete namespace %s: ", ns, utils.Err(err)))
	}
}
