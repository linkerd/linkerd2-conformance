package inject

import (
	"fmt"

	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/linkerd/linkerd2/testutil"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
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
		gomega.Expect(1).To(gomega.Equal(2))
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
