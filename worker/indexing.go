package worker

import (
	. "../shared"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/esapi"
	"path/filepath"
	"strings"
)

type IndexedFile struct {
	Filename string `json:"filename"`
	Repository string `json:"repository"`
	Language string `json:"language"`
	Lines []string `json:"lines"`
}

type IndexedReport struct {
	Repository string `json:"repository"`
	Report Reports `json:"report"`
}

func IndexReport(repo string, report Reports){

	doc := IndexedReport{
		Repository: repo,
		Report:     report,
	}

	indexDocument("report", RenderJSON(doc))
}

func IndexCodeFile(repo string, file string, content string){
	programingLang := ProgrammingLanguages[filepath.Ext(file)]
	if len(programingLang) == 0 { return }

	doc := IndexedFile{
		Filename:   file,
		Repository: repo,
		Language: programingLang,
		Lines:      splitLines(content),
	}

	indexDocument("files", RenderJSON(doc))
}

func ClearRepositoryFromIndex(repository string){

	deleteReq := esapi.DeleteByQueryRequest{
		Index:               []string{"files", "report"},
		Query:               "\""+repository+"\"",

	}

	res, err := deleteReq.Do(context.Background(), esClient)
	CheckError(err)
	res.Body.Close()
}

func indexDocument(index string, document string){

	req := esapi.IndexRequest{
		Index:               	index,
		Body:                	strings.NewReader(document),
		Refresh:             	"true",
	}

	res, err := req.Do(context.Background(), esClient)
	CheckError(err)
	res.Body.Close()
}

func splitLines(s string) []string{
	var col []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		col = append(col, scanner.Text())
	}
	return col
}


func RenderJSON(str interface{}) string {
	bytes, err := json.Marshal(str)

	if err != nil {
		fmt.Println(err)
	}

	return string(bytes)
}