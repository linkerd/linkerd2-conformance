package utils

import (
	"github.com/linkerd/linkerd2/testutil"
)

var TestHelper *testutil.TestHelper

func InitTestHelper() {
	TestHelper = testutil.NewTestHelper()
}
