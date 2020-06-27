package specs

import (
	"testing"

	"github.com/linkerd/linkerd2-conformance/specs/inject"
	"github.com/linkerd/linkerd2-conformance/specs/install"
	"github.com/linkerd/linkerd2-conformance/specs/uninstall"
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func RunAllSpecs(t *testing.T) {

	h := utils.TestHelper
	c := utils.TestConfig

	globalControlPlane := c.GlobalControlPlane()
	u := h.Uninstall()
	ha := c.HA()

	// A single top-level wrapper Describe is required to prevent
	// the specs from being run in a random order
	// The Describe message is intentionally left empty
	// as it only serves to prevent randomization of specs
	_ = ginkgo.Describe("", func() {

		// Run Install / Uninstall test in a separate
		// Describe block before running the primary tests
		// This is done so that the BeforeEach and AfterEach
		// blocks do not interfere with these tests
		_ = ginkgo.Describe("", func() {
			_ = install.RunInstallSpec()
			if !globalControlPlane { // Immediately uninstall if each test shall have its own control-plane
				_ = uninstall.RunUninstallSpec()
			}
		})

		// Run primary tests
		_ = ginkgo.Describe("", func() {

			if !globalControlPlane {
				_ = ginkgo.BeforeEach(func() {
					utils.InstallLinkerdControlPlane(h, ha)
				})

				_ = ginkgo.AfterEach(func() {
					utils.UninstallLinkerdControlPlane(h)
				})

			}

			// Bring main tests into scope
			_ = inject.RunInjectSpec()

			// global uninstall (if true) should always run at the end
			if globalControlPlane && u {
				_ = uninstall.RunUninstallSpec()
			}
		})
	})

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "linkerd2 conformance validation")
}
