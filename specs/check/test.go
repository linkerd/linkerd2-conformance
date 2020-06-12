package check

import (
	"encoding/json"
	"fmt"

	"github.com/linkerd/linkerd2/testutil"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

type CheckOutput struct {
	Success    bool `json:"success"`
	Categories []struct {
		CategoryName string `json:"categoryName"`
		Checks       []struct {
			Result string `json:"result"`
			Error  string `json:"error"`
		}
	}
}

func getFailedChecks(r *CheckOutput) string {
	err := "The following errors were detected:\n"

	for _, c := range r.Categories {
		for _, check := range c.Checks {
			if check.Result == "error" {
				err = fmt.Sprintf("%s\n%s", err, check.Error)
			}
		}
	}

	return err
}

func testCheck(h *testutil.TestHelper, pre bool) {
	var checkResult *CheckOutput

	cmd := []string{
		"check",
		// "--expected-version",
		// h.GetVersion(),
		"-o",
		"json",
	}

	if pre {
		cmd = append(cmd, "--pre")
	}

	ginkgo.By("running check command")
	out, stderr, err := h.LinkerdRun(cmd...)
	gomega.Expect(stderr).To(gomega.Equal(""))

	ginkgo.By("validating `check` output")
	err = json.Unmarshal([]byte(out), &checkResult)
	gomega.Expect(err).To(gomega.BeNil())

	gomega.Expect(checkResult.Success).Should(gomega.BeTrue(), getFailedChecks(checkResult))
}
