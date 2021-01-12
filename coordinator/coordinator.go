package coordinator

import (
	. "../shared"
)

func RunCoordinator() {
	Log("Initializing Coordinator")
	go ControlServer()
	go InitStatusSpace(StatusSpace())
	go InitTaskDelegator(TaskSpace())
	<- make(chan string)
}


