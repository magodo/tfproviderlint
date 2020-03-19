package commentignore

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"golang.org/x/tools/go/analysis"
)

func TestValidateAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, testAnalyzer, "a")
}

var testAnalyzer = &analysis.Analyzer{
	Name: "commentignore",
	Doc:  "find ignore comments for later passes",
	Run:  testrun,
}

func testrun(pass *analysis.Pass) (interface{}, error) {
	ignoreMap := collectIgnores(pass)
	for ruleid, ignores := range ignoreMap {
		for _, ignore := range ignores {
			pass.Reportf(ignore.Pos, "%s ignored", ruleid)
		}
	}
	return nil, nil
}
