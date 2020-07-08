package ingress

import (
	"testing"

	"github.com/linkerd/linkerd2-conformance/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func() {
	utils.BeforeSuiteCallback()
	utils.TestEmojivotoApp()
	utils.TestEmojivotoInject()
})

func TestIngress(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ingress")
}

var _ = AfterSuite(func() {
	utils.AfterSuiteCallBack()
	utils.TestEmojivotoUninstall()
})
