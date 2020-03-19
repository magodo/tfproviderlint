package commentignore

import (
	"go/ast"
	"go/token"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const commentIgnorePrefix = "lintignore:"

var Analyzer = &analysis.Analyzer{
	Name:       "commentignore",
	Doc:        "find ignore comments for later passes",
	Run:        run,
	ResultType: reflect.TypeOf(new(Ignorer)),
}

type ignore struct {
	Pos token.Pos
	End token.Pos
}

type Ignorer struct {
	ignores map[string][]ignore
}

func (ignorer *Ignorer) ShouldIgnore(key string, n ast.Node) bool {
	for _, ig := range ignorer.ignores[key] {
		if ig.Pos <= n.Pos() && ig.End >= n.End() {
			return true
		}
	}

	return false
}

func collectIgnores(pass *analysis.Pass) map[string][]ignore {
	ignores := map[string][]ignore{}
	for _, f := range pass.Files {
		cmap := ast.NewCommentMap(pass.Fset, f, f.Comments)
		for n, cgs := range cmap {
			for _, cg := range cgs {
				if strings.HasPrefix(cg.Text(), commentIgnorePrefix) {
					key := strings.TrimPrefix(cg.Text(), commentIgnorePrefix)
					key = strings.TrimSpace(key)

					// is it possible for nested pos/end to be outside the largest nodes?
					ignores[key] = append(ignores[key], ignore{n.Pos(), n.End()})
				}
			}
		}
	}
	return ignores
}

func run(pass *analysis.Pass) (interface{}, error) {
	return &Ignorer{
		ignores: collectIgnores(pass),
	}, nil
}
