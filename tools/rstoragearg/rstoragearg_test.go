package ctxarg

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestRStorageArg(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer)
}
