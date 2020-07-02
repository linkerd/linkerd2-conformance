package <testname>

import (
	"testing"

	"github.com/linkerd/linkerd2-conformance/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(utils.BeforeSuiteCallback)

func Test<testname>(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "<testname>")
}

var _ = AfterSuite(utils.AfterSuiteCallBack)
