package conformance_test

import (
	"fmt"
	"testing"

	"github.com/linkerd/linkerd2-conformance/specs"
	"github.com/linkerd/linkerd2-conformance/utils"
)

func TestMain(m *testing.M) {
	err := utils.InitTestHelper()
	if err != nil {
		fmt.Println(err)
	}
	_ = m.Run()
}

func TestConformance(t *testing.T) {
	specs.RunAllSpecs(t)
}
