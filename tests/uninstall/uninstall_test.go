package uninstall

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUnistall(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uninstall")
}
