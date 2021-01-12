package coordinator

import (
	. "../shared"
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"io"
)

var (
	statusSpace space.Space
)

func InitStatusSpace(){
	uri := StatusSpaceUri()
	statusSpace = space.NewSpace(uri)
	Log(fmt.Sprintf("Status space initialized at: %v", uri))
}

func InitLogsSpace(){
	uri := LogSpaceUri()
	LogSpace = space.NewSpace(uri)
	Log(fmt.Sprintf("log space initialized at: %v", uri))
}

func GetStatusReports(writer io.Writer){
	var workerId string
	var status string
	var since string

	query, _ := statusSpace.QueryAll(&workerId, &status, &since)
	for _, t := range query {
		writer.Write([]byte(t.String()+"\n"))
	}
}

func GetLogs(writer io.Writer){
	var s string
	query, _ := LogSpace.QueryAll(&s)
	for _, t := range query {
		writer.Write([]byte(t.GetFieldAt(0).(string)))
	}
}

func GetTasks(writer io.Writer){
	var s string
	query, _ := taskSpace.QueryAll(&s)
	for _, t := range query {
		writer.Write([]byte(t.GetFieldAt(0).(string)+"\n"))
	}
}

func GetRepositories(writer io.Writer){
	var s string
	query, _ := repositoriesSpace.QueryAll(&s, &s)
	for _, t := range query {
		writer.Write([]byte(t.GetFieldAt(0).(string)+ " " +t.GetFieldAt(1).(string)+"\n"))
	}
}

