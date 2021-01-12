package worker

import (
	. "../shared"
	"github.com/DanielHauge/goSpace/space"
	"time"
)

var taskSpace space.Space

func InitTaskReceiver(){

	taskSpace = space.NewRemoteSpace(TaskSpace())
	Logf("Worker: %v connected to task space", workerId)

	var repo string
	for {
		taskSpace.Get(&repo)
		ReportTaskBegin(repo)

		// Do work on repo.
		<- time.After(time.Second*3)
		// Could do something like a subtask space, with name of the task, and goroutines to take tasks, and


		ReportTaskDone(repo)
	}
}
