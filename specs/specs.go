package specs

import (
	"testing"

	"github.com/linkerd/linkerd2-conformance/specs/inject"
	"github.com/linkerd/linkerd2-conformance/utils"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func RunAllSpecs(t *testing.T) {

	c := utils.TestConfig
	h := utils.TestHelper

	// Install a single global control plane
	if c.GlobalControlPlane.Enable {
		_ = ginkgo.BeforeSuite(func() {
			utils.InstallLinkerdControlPlane(h)
		})

		if c.GlobalControlPlane.Uninstall {
			_ = ginkgo.AfterSuite(func() {
				utils.UninstallLinkerdControlPlane(h)
			})
		}
	} else {
		// Install and uninstall a control plane
		// before and after each It block
		// This means, while writing a new test suite,
		// each major feature must be wrapped in its own
		// It block
		_ = ginkgo.BeforeEach(func() {
			utils.InstallLinkerdControlPlane(h)
		})

		_ = ginkgo.AfterEach(func() {
			utils.UninstallLinkerdControlPlane(h)
		})

	}

	// A single top-level wrapper Describe is required to prevent
	// the specs from being run in a random order
	// The Describe message is intentionally left empty
	// as it only serves to prevent randomization of specs
	_ = ginkgo.Describe("", func() {

		// Bring tests into scope
		_ = inject.RunInjectSpec()
	})

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "linkerd2 conformance validation")
}
