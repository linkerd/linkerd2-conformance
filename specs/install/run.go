package install

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

// RunInstallSpec runs the install specs
func RunInstallSpec(t *testing.T, desc string) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, desc)
}
