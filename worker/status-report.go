package worker

import (
	. "../shared"
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"time"
)


var statusSpace space.Space


func ConnectWorkerStatus(){
	statusSpace = space.NewRemoteSpace(StatusSpace())
	statusSpace.Put(workerId, "idle", time.Now())
	Logf("Worker: %v connected to status space", workerId)
}

func CloseStatusSpace(){
	var s string
	var t time.Time
	statusSpace.GetP(workerId, &s, &t)
	Logf("Worker: %v disconnected from status space", workerId)
}

func ReportTaskBegin(task string){
	var s string
	var t time.Time
	statusSpace.Get(workerId, &s, &t)
	statusSpace.Put(workerId, fmt.Sprintf("Working on: %v", task), time.Now())
	Logf("Worker: %v began working on: %v", workerId, task)
}

func ReportTaskDone(task string){
	var s string
	var t time.Time
	statusSpace.Get(workerId, &s, &t)
	statusSpace.Put(workerId, fmt.Sprintf("Done with: %v", task), time.Now())
	Logf("Worker: %v done with: %v", workerId, task)
}