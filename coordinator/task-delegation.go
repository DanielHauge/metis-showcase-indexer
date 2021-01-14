package coordinator

import (
	. "../shared"
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"github.com/robfig/cron"
)

var repositoriesSpace space.Space

func InitIndexDelegator(){
	uri := IndexSpaceUri()
	IndexSpace = space.NewSpace(uri)
	repositoriesSpace = space.NewSpace("repo")

	Log(fmt.Sprintf("Task space initialized at: %v", uri))


	var r string
	repositoriesSpace.Get(&r)
	for {
		Log(fmt.Sprintf("Repository to be reindex: %v", r))
		IndexSpace.Put(r)
		c := cron.New()
		c.AddFunc("0 */12 * * *", func() { repositoriesSpace.PutP(r) })
		c.Start()
		repositoriesSpace.Get(&r)
	}


}

