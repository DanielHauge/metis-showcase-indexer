package coordinator

import (
	. "../shared"
	"fmt"
	"github.com/DanielHauge/goSpace/space"
)

func RunCoordinator() {
	ClientName = "Coordinator"
	InitLogsSpace()
	go ControlServer()
	go InitStatusSpace()
	go InitIndexDelegator()
	go InitTaskSpace()
	Log("Fully Initializing")
	<- make(chan string)
}

func InitStatusSpace(){
	uri := StatusSpaceUri()
	StatusSpace = space.NewSpace(uri)
	Log(fmt.Sprintf("Status space initialized at: %v", uri))
}

func InitLogsSpace(){
	uri := LogSpaceUri()
	LogSpace = space.NewSpace(uri)
	Log(fmt.Sprintf("log space initialized at: %v", uri))
}

func InitTaskSpace(){
	uri := TaskSpaceUri()
	TaskSpace = space.NewSpace(uri)
	Log(fmt.Sprintf("Task space initialized at: %v", uri))
}


