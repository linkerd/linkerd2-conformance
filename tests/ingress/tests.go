package ingress

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func pingEmojivoto(ip string) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s", ip), nil)
	if err != nil {
		return err
	}

	req.Host = "example.com"

	client := http.Client{
		Timeout: time.Minute,
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("did not recieve status code 200. Recieved %d", res.StatusCode)
	}
	return nil

}

func testNginx() {
	h, _ := utils.GetHelperAndConfig()
	ginkgo.By("Creating ingress-nginx controller")
	_, err := h.Kubectl("", "apply", "-f", "testdata/controllers/nginx.yaml")

	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to create controller: %s", utils.Err(err)))

	err = h.CheckPods(utils.NginxNs, utils.NginxController, 1)
	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to verify controller pods: %s", utils.Err(err)))

	ginkgo.By("Injecting linkerd into the ingress controller pods")
	out, err := h.Kubectl("", "get", "-n", utils.NginxNs, "deploy", utils.NginxController, "-o", "yaml")
	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to get YAML manifest for deploy/%s: %s", utils.NginxController, utils.Err(err)))

	out, stderr, err := h.PipeToLinkerdRun(out, "inject", "-")
	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to inject: %s", stderr))

	_, err = h.KubectlApply(out, utils.NginxNs)
	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to apply injected manifests: %s", utils.Err(err)))

	err = h.CheckPods(utils.NginxNs, utils.NginxController, 1)
	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to verify controller pods: %s", utils.Err(err)))

	ginkgo.By("Verifying if ingress controller pods have been injected")
	// Wait upto 3mins for proxy container to show up
	err = utils.CheckProxyContainer(utils.NginxController, utils.NginxNs)
	gomega.Expect(err).Should(gomega.BeNil(), utils.Err(err))

	ginkgo.By("Applying ingress resource")
	_, err = h.Kubectl("", "apply", "-f", "testdata/resources/nginx.yaml")
	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to create ingress resource: %s", utils.Err(err)))

	ginkgo.By("Checking if emojivoto is reachable")
	ip, err := h.Kubectl("", "get", "svc", "-n", utils.NginxNs, utils.NginxController, "-o", "jsonpath='{.status.loadBalancer.ingress[0].ip}'")
	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to get external Ip address for service/%s: %s", utils.NginxController, utils.Err(err)))
	log.Println(ip)
	err = pingEmojivoto(strings.Trim(ip, "'"))
	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to reach emojivoto: %s", utils.Err(err)))

	ginkgo.By(fmt.Sprintf("Removing ingress controller in namespace %s", utils.NginxNs))
	_, err = h.Kubectl("", "delete", "ns", utils.NginxNs)

	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to delete resources in namespace %s", utils.NginxNs))
}
