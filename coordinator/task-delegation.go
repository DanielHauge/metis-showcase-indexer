package coordinator

import (
	. "../shared"
	"github.com/DanielHauge/goSpace/space"
	"time"
)

var repositoriesSpace space.Space

func InitTaskDelegator(uri string){

	taskSpace := space.NewSpace(uri)
	repositoriesSpace = space.NewSpace("repo")

	Log("Task delegation space and repository space initialized")

	var r string
	var t time.Time
	for {
		repositoriesSpace.Get(&r, &t)
		Logf("Repository to be reindex: %v at %v", r, t)
		now := time.Now()
		if t.After(now) {
			until := t.Sub(now)
			Logf("Waiting with task delegation for %v", until)
			<- time.After(until)
		}
		taskSpace.Put(r)
		Logf("Repository: %v got delegated", r)
		repositoriesSpace.Put(r, time.Now().Add(time.Second * 10))
	}
}