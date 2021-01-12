package coordinator

import (
	. "../shared"
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"time"
)

var repositoriesSpace space.Space
var taskSpace space.Space


func InitTaskDelegator(){
	uri := TaskSpaceUri()
	taskSpace = space.NewSpace(uri)
	repositoriesSpace = space.NewSpace("repo")

	Log(fmt.Sprintf("Task space initialized at: %v", uri))

	var r string
	var tStr string
	repositoriesSpace.Get(&r, &tStr)
	for {
		Log(fmt.Sprintf("Repository to be reindex: %v at %v", r, tStr))
		now := time.Now()
		t, _ := time.Parse(TimeFormat, tStr)
		if t.After(now) {
			until := t.Sub(now)
			Log(fmt.Sprintf("Waiting with task delegation for %v", until))
			<- time.After(until)
		}
		taskSpace.Put(r)
		Log(fmt.Sprintf("Repository: %v got delegated", r))
		repositoriesSpace.PutP(r, time.Now().Add(time.Second * 10).Format(TimeFormat))
		repositoriesSpace.Get(&r, &tStr)
	}
}

