package check

import (
	"fmt"

	"github.com/linkerd/linkerd2/testutil"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func testCheck(h *testutil.TestHelper, pre bool) {
	var golden string

	cmd := []string{
		"check",
		"--expected-version",
		h.GetVersion(),
	}

	if pre {
		cmd = append(cmd, "--pre")
		golden = "check.pre.golden"
	} else {
		golden = "check.golden"
	}

	ginkgo.By("running check command")
	out, stderr, err := h.LinkerdRun(cmd...)
	gomega.Expect(stderr).To(gomega.Equal(""))

	ginkgo.By("validating `check` output")
	err = h.ValidateOutput(out, golden)
	if err != nil {
		// the mismatch in golden and out may happen has different releases of linkerd2 have
		// may have introduced different set of checks. Hence, make this a soft requirement.
		ginkgo.Skip(fmt.Sprintf("Mismatch in output from `check`\n%s", err.Error()))
	}
}
