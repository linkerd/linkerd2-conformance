package install

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/linkerd/linkerd2/testutil"
	"github.com/onsi/ginkgo"
)

// NewInstallSpec returns a new install test spec
func NewInstallSpec() bool {
	return ginkgo.Context("`linkerd install`", func() {
		var h *testutil.TestHelper

		ginkgo.BeforeEach(func() {
			h = utils.TestHelper
		})

		// Add test cases here

		ginkgo.It("should successfully install a Linkerd control plane", func() {
			testControlPlaneInstall(h)
			testControlPlaneState(h)
		})

	})
}
