package worker

import (
	. "../shared"
	"fmt"
	"github.com/beevik/guid"
)



func RunWorker(){
	workerId := guid.New().String()
	ClientName = fmt.Sprintf("Worker: %v", workerId)
	ConnectLogSpace()
	ConnectTaskSpace()
	defer ConnectStatusSpace(workerId)()

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
		case INDEXFILE:
			go IndexFile(repository, fileName, fileContent)
			TaskSpace.Put(managerId, RESULT, task, repository, "indexed")
		case ANALYSE_GO:
			report := AnalyseGoFile(fileName, fileContent)
			TaskSpace.Put(managerId, RESULT, task, repository, report.RenderJSON())
		default:
			TaskSpace.Put(managerId, RESULT, task, repository, fmt.Sprintf("error:%v is not a valid task", task))
		}
		ReportDone(taskDescription)

	}
}

