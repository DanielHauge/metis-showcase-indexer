package shared

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"math"
)

// Inspiration from: https://github.com/droptheplot/abcgo/blob/master/main.go

type Reports []Report

type Report struct {
	Path       string `json:"path"`
	Line       int    `json:"line"`
	Name       string `json:"name"`
	Assignment int    `json:"assignment"`
	Branch     int    `json:"branch"`
	Condition  int    `json:"condition"`
	Score      int    `json:"score"`
}

func ReportNode(report *Report, n ast.Node) {
	switch n := n.(type) {
	case *ast.AssignStmt, *ast.IncDecStmt:
		report.Assignment++
	case *ast.CallExpr:
		report.Branch++
	case *ast.IfStmt:
		if n.Else != nil {
			report.Condition++
		}
	case *ast.BinaryExpr, *ast.CaseClause:
		report.Condition++
	}
}

func (report *Report) Calc() {
	a := math.Pow(float64(report.Assignment), 2)
	b := math.Pow(float64(report.Branch), 2)
	c := math.Pow(float64(report.Condition), 2)

	report.Score = int(math.Sqrt(a + b + c))
}

func (reports Reports) RenderJSON() string {
	bytes, err := json.Marshal(reports)

	if err != nil {
		fmt.Println(err)
	}

	return string(bytes)
}

func ParseReports(jsonStr string) Reports {
	var report Reports

	e := json.Unmarshal([]byte(jsonStr), &report)
	if e != nil {
		fmt.Println(e)
	}
	return report
}
