package worker

import (
	. "../shared"
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"time"
)


var statusSpace space.Space


func ConnectWorkerStatus(){
	uri := StatusSpaceUri()
	statusSpace = space.NewRemoteSpace(uri)
	statusSpace.Put(workerId, "idle", time.Now().Format(TimeFormat))
	Log(fmt.Sprintf("Connected to status space at: %v", uri))
}

func ConnectLogSpace(){
	uri := LogSpaceUri()
	LogSpace = space.NewRemoteSpace(uri)
	Log(fmt.Sprintf("Connected to log space at: %v", uri))
}

func CloseStatusSpace(){
	var s string
	var t string
	statusSpace.GetP(workerId, &s, &t)
	Log("Disconnected from status space")
}

func ReportTaskBegin(task string){
	var s string
	var t string
	statusSpace.Get(workerId, &s, &t)
	workStr := fmt.Sprintf("Working on: %v", task)
	statusSpace.Put(workerId, workStr, time.Now().Format(TimeFormat))
	Log(workStr)
}

func ReportTaskDone(task string){
	var s string
	var t string
	statusSpace.Get(workerId, &s, &t)
	doneStr := fmt.Sprintf("Done with: %v", task)
	statusSpace.Put(workerId, doneStr, time.Now().Format(TimeFormat))
	Log(doneStr)
}