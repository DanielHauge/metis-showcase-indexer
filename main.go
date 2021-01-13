package main

import (
	. "./coordinator"
	. "./manager"
	. "./shared"
	. "./worker"
	"errors"
	"os"
)

func main() {
	defer LogPanicError()
	mode := os.Getenv("mode")
	switch mode {
		case Manager: RunManager()
		case Coordinator: RunCoordinator()
		case Worker: RunWorker()
		case "Test": RunTesting()
		default: panic(errors.New("No default mode found, please set environment variable: 'mode' like -e mode=manager"))
	}
}

func RunTesting(){
	/*
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/DanielHauge/metis-showcase-indexer",
	})
	if err != nil{
		Log("Failed with: " + err.Error() )
	}

	ref, err := repo.Head()
	headCommit, err := repo.CommitObject(ref.Hash())
	files, err := headCommit.Files()
	files.ForEach(func(file *object.File) error {
		fmt.Println(file.Name)
		content, _ := file.Contents()
		fmt.Println(content)
		return nil
	})
	*/

	/*
	var src = "package worker\n\nimport (\n\t\"go/ast\"\n\t\"go/parser\"\n\t\"go/token\"\n\t\"math\"\n)\n\n// Inspiration from: https://github.com/droptheplot/abcgo/blob/master/main.go\n\n// Reports is a collection of Report.\ntype Reports []Report\n\n// Report contains statistics for single function.\ntype Report struct {\n\tPath       string `json:\"path\"`\n\tLine       int    `json:\"line\"`\n\tName       string `json:\"name\"`\n\tAssignment int    `json:\"assignment\"`\n\tBranch     int    `json:\"branch\"`\n\tCondition  int    `json:\"condition\"`\n\tScore      int    `json:\"score\"`\n}\n\nfunc AnalyseGoFile(content string) Reports{\n\n\n\tfileSet := token.NewFileSet()\n\n\tnode, _ := parser.ParseFile(fileSet, \"\", content, 0)\n\treports := reportFile(fileSet, node)\n\treturn reports\n}\n\nfunc reportFile(fset *token.FileSet, n ast.Node) Reports {\n\tvar reports Reports\n\n\tast.Inspect(n, func(n ast.Node) bool {\n\t\tif fn, ok := n.(*ast.FuncDecl); ok {\n\t\t\treport := Report{\n\t\t\t\tPath: fset.File(fn.Pos()).Name(),\n\t\t\t\tLine: fset.Position(fn.Pos()).Line,\n\t\t\t\tName: fn.Name.Name,\n\t\t\t}\n\n\t\t\tast.Inspect(n, func(n ast.Node) bool {\n\t\t\t\treportNode(&report, n)\n\t\t\t\treturn true\n\t\t\t})\n\n\t\t\treport.Calc()\n\t\t\treports = append(reports, report)\n\t\t\treturn false\n\t\t}\n\t\treturn true\n\t})\n\n\treturn reports\n}\n\nfunc reportNode(report *Report, n ast.Node) {\n\tswitch n := n.(type) {\n\tcase *ast.AssignStmt, *ast.IncDecStmt:\n\t\treport.Assignment++\n\tcase *ast.CallExpr:\n\t\treport.Branch++\n\tcase *ast.IfStmt:\n\t\tif n.Else != nil {\n\t\t\treport.Condition++\n\t\t}\n\tcase *ast.BinaryExpr, *ast.CaseClause:\n\t\treport.Condition++\n\t}\n}\n\nfunc (report *Report) Calc() {\n\ta := math.Pow(float64(report.Assignment), 2)\n\tb := math.Pow(float64(report.Branch), 2)\n\tc := math.Pow(float64(report.Condition), 2)\n\n\treport.Score = int(math.Sqrt(a + b + c))\n}"
	reports := AnalyseGoFile("gg.go",src)
	reportJson := reports.RenderJSON()
	fmt.Println(reportJson)
	_ = reports
	*/

}

