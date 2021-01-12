package coordinator

import (
	. "../shared"
)

func RunCoordinator() {
	ClientName = "Coordinator"
	InitLogsSpace()
	go ControlServer()
	go InitStatusSpace()
	go InitTaskDelegator()
	Log("Fully Initializing")
	<- make(chan string)
}


