package shared

import (
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"os"
	"time"
)

var StatusId string

func StatusSpaceUri() string{
	return "tcp://" + withDefault(os.Getenv("host"), "localhost") + ":9092/status"
}

func ConnectStatusSpace(id string) func(){
	StatusId = id
	uri := StatusSpaceUri()
	StatusSpace = space.NewRemoteSpace(uri)
	StatusSpace.Put(StatusId, "idle", time.Now().Format(TimeFormat))
	Log(fmt.Sprintf("Connected to status space at: %v", uri))
	return CloseStatusSpace
}

func CloseStatusSpace(){
	var s string
	var t string
	StatusSpace.GetP(StatusId, &s, &t)
	Log("Disconnected from status space")
}

func ReportStarted(task string){
	var s string
	var t string
	StatusSpace.Get(StatusId, &s, &t)
	message := fmt.Sprintf("Working on: %v", task)
	StatusSpace.Put(StatusId, message, time.Now().Format(TimeFormat))
	Log(message)
}

func ReportDone(task string){
	var s string
	var t string
	StatusSpace.Get(StatusId, &s, &t)
	message := fmt.Sprintf("Done with: %v", task)
	StatusSpace.Put(StatusId, message, time.Now().Format(TimeFormat))
	Log(message)
}