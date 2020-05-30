package conformance_test

import (
	"testing"

	"github.com/linkerd/linkerd2-conformance/specs/install"
	"github.com/linkerd/linkerd2-conformance/utils"
)

type runner func(*testing.T, string)

func TestMain(m *testing.M) {
	utils.InitTestHelper()
	_ = m.Run()
}

func TestConformance(t *testing.T) {

	specs := []struct {
		run  runner
		desc string
	}{
		{
			run:  install.RunInstallSpec,
			desc: "control plane installation spec",
		},
	}

	for _, s := range specs {
		s.run(t, s.desc)
	}
}
