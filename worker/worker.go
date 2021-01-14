package worker

import (
	. "../shared"
	"fmt"
	"github.com/beevik/guid"
	"github.com/elastic/go-elasticsearch"
)

var esClient *elasticsearch.Client

func RunWorker(){
	InitUtility()
	workerId := guid.New().String()
	ClientName = fmt.Sprintf("Worker: %v", workerId)
	ConnectLogSpace()
	ConnectTaskSpace()
	defer ConnectStatusSpace(workerId)()

	c, err := elasticsearch.NewDefaultClient()
	CheckErrorF(err)
	esClient = c

	var managerId string
	var task string
	var repository string
	var fileName string
	var fileContent string

	for {
		if _, e := TaskSpace.Get(&managerId, TASK, &task, &repository, &fileName, &fileContent); CheckError(e) { continue }
		taskDescription := fmt.Sprintf("%v for %v", task, fileName)
		ReportStarted(taskDescription)
		Log(taskDescription)
		switch task {
		case INDEX_FILE:
			go IndexCodeFile(repository, fileName, fileContent)
			_,e := TaskSpace.Put(managerId, RESULT, task, repository, "indexed")
			CheckError(e)
		case ANALYSE_GO:
			report := AnalyseGoFile(fileName, fileContent)
			_,e := TaskSpace.Put(managerId, RESULT, task, repository, report.RenderJSON())
			CheckError(e)
		case INDEX_REPORT:
			IndexReport(repository, ParseReports(fileContent))
		case CLEAR:
			go ClearRepositoryFromIndex(repository)
		default:
			_,e := TaskSpace.Put(managerId, RESULT, task, repository, fmt.Sprintf("error:%v is not a valid task", task))
			CheckError(e)
		}
		ReportDone(taskDescription)

	}
}

