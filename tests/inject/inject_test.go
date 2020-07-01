package inject

import (
	"testing"

	"github.com/linkerd/linkerd2-conformance/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(utils.BeforeSuiteCallback)

func TestInject(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Inject")
}

var _ = AfterSuite(utils.AfterSuiteCallBack)
