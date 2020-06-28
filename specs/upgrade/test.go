package upgrade

import (
	"fmt"

	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func testUpgradeCLI() {
	h, c := utils.GetHelperAndConfig()

	ginkgo.By(fmt.Sprintf("Upgrading CLI from version %s to %s", h.UpgradeFromVersion(), h.GetVersion()))

	err := utils.InstallLinkerdBinary(c.GetLinkerdPath(), h.GetVersion(), true, false)
	gomega.Expect(err).Should(gomega.BeNil(), utils.Err(err))

	cmd := []string{
		"version",
		"--short",
		"--client",
	}

	ginkgo.By("Validating CLI version")
	out, stderr, err := h.LinkerdRun(cmd...)
	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("could not run `linkerd version command`: %s", stderr))
	gomega.Expect(out).Should(gomega.ContainSubstring(h.GetVersion()), "failed to upgrade CLI")
}

func testUpgrade() {
	h, _ := utils.GetHelperAndConfig()

	cmd := "upgrade"

	ginkgo.By("Running `linkerd upgrade` command")
	out, stderr, err := h.LinkerdRun(cmd)
	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("`linkerd upgrade` command failed: %s", stderr))

	_, err = h.Kubectl(out, "apply", "--prune", "-l", "linkerd.io/control-plane-ns="+h.GetLinkerdNamespace(), "-f", "-")

	gomega.Expect(err).Should(gomega.BeNil(), fmt.Sprintf("failed to apply manifests: %s", utils.Err(err)))

	// TODO: Once https://github.com/linkerd/linkerd2/pull/4681 is merged, check the state of control plane deployments

	utils.RunCheck(h, false)
}
