package rstoragearg

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = `check for parameters order while receiving RemoteStorage as parameter

The RemoteStorage checker walks functions checking RemoteStorage parameters 
and ensures the RemoteStorage parameter is always the first received argument.`

var Analyzer = &analysis.Analyzer{
	Name:     "rstoragearg",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

var name, paramorder *bool

func init() {
	name = Analyzer.Flags.Bool("name", true, "will ensure RemoteStorage as parameter is named rstorage")
	paramorder = Analyzer.Flags.Bool("paramorder", true, "will ensure rstorage is the first argument of the functions")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	typeName := `cobrago.RemoteStorage`
	expectedVariableName := `rstorage`

	nodeFilter := []ast.Node{
		(*ast.FuncType)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		function := n.(*ast.FuncType)

		numRStorage := 0
		for _, f := range function.Params.List {
			t := pass.TypesInfo.TypeOf(f.Type)
			if typeName == t.String() || strings.HasSuffix(t.String(), typeName) {
				numRStorage++
			}
		}

		if numRStorage > 1 {
			pass.Reportf(function.Pos(), "the function has more than one RemoteStorage defined in the parameters.")
			return
		}

		for key, f := range function.Params.List {
			t := pass.TypesInfo.TypeOf(f.Type)
			if !(typeName == t.String() || strings.HasSuffix(t.String(), typeName)) {
				continue
			}

			if *paramorder {
				if key != 0 {
					pass.Reportf(function.Pos(), "the function has a RemoteStorage that is not the first argument.")
				}
			}
			if *name {
				if len(f.Names) > 1 || len(f.Names) == 0 {
					continue
				}
				if f.Names[0].Name != expectedVariableName {
					pass.Reportf(function.Pos(), fmt.Sprintf("the function has a RemoteStorage that is not named '%s'.", expectedVariableName))
				}
			}
		}
	})
	return nil, nil
}
