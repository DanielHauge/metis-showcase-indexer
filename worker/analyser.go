package worker

import (
	. "../shared"
	"go/ast"
	"go/parser"
	"go/token"
)

// Inspiration from: https://github.com/droptheplot/abcgo/blob/master/main.go


func AnalyseGoFile(name string, content string) Reports{

	fileSet := token.NewFileSet()
	node, _ := parser.ParseFile(fileSet, name, content, 0)
	reports := reportFile(fileSet, node)
	return reports
}

func reportFile(fset *token.FileSet, n ast.Node) Reports {
	var reports Reports

	ast.Inspect(n, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			report := Report{
				Path: fset.File(fn.Pos()).Name(),
				Line: fset.Position(fn.Pos()).Line,
				Name: fn.Name.Name,
			}

			ast.Inspect(n, func(n ast.Node) bool {
				ReportNode(&report, n)
				return true
			})

			report.Calc()
			reports = append(reports, report)
			return false
		}
		return true
	})

	return reports
}

