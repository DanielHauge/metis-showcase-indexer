package worker

import (
	. "../shared"
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"time"
)

var taskSpace space.Space

func InitTaskReceiver(){

	taskSpaceUri := TaskSpaceUri()
	taskSpace = space.NewRemoteSpace(taskSpaceUri)
	Log(fmt.Sprintf("Connected to task space at: %v", taskSpaceUri))

	var r string

	for {
		taskSpace.Get(&r)
		ReportTaskBegin(r)
		// Do work on repo.
		<- time.After(time.Second*3)
		// Could do something like a subtask space, with name of the task, and goroutines to take tasks, and

		ReportTaskDone(r)
	}
}
