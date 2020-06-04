package install

import (
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/linkerd/linkerd2/testutil"
	"github.com/onsi/ginkgo"
)

var _ = ginkgo.Describe("linkerd2 control plane components", func() {
	var h *testutil.TestHelper

	ginkgo.BeforeEach(func() {
		h = utils.TestHelper
	})

	// Add test cases here

	ginkgo.It("should successfully successfully install", func() {
		testControlPlaneInstall(h)
	})

	ginkgo.It("should be in a running state", func() {
		testControlPlaneState(h)
	})

})
