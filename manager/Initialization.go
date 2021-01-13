package manager

import (
	. "../shared"
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing/object"
	"github.com/go-git/go-git/storage/memory"
	"strings"
)



func InitManager(){

	indexSpaceUri := IndexSpaceUri()
	taskSpaceUri := TaskSpaceUri()
	IndexSpace = space.NewRemoteSpace(indexSpaceUri)
	TaskSpace = space.NewRemoteSpace(taskSpaceUri)
	Log(fmt.Sprintf("Connected to index space at: %v", indexSpaceUri))
	Log(fmt.Sprintf("Connected to task space at: %v", taskSpaceUri))

	var r string

	for {
		IndexSpace.Get(&r)
		ReportStarted(r)

		repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{	URL: r	})
		if CheckError(err) { continue }
		ref, err := repo.Head()
		CheckErrorF(err)
		headCommit, err := repo.CommitObject(ref.Hash())
		CheckErrorF(err)
		files, err := headCommit.Files()
		CheckErrorF(err)

		var tasksCount int
		var goReports []Report
		tasksCount = 0
		files.ForEach(func(file *object.File) error {
			if file == nil{ return nil}
			tasksCount++
			n := file.Name

			content, e := file.Contents()

			if !CheckError(e){
				TaskSpace.Put(ClientName, TASK, INDEXFILE, r, n, content )
				if strings.HasSuffix(n, ".go") {
					TaskSpace.Put(ClientName, TASK, ANALYSE_GO, r, n, content )
				}
				if strings.HasSuffix(n, ".showcase") {
					TaskSpace.Put(ClientName, TASK, SHOWCASE, r, n, content )
				}
			}
			return nil
		})

		var res string
		for tasksCount > 0 {
			var taskName string
			TaskSpace.Get(ClientName, RESULT, &taskName, r, &res)
			switch taskName {
			case ANALYSE_GO:
				reports := ParseReports(res)

				goReports = append(goReports, reports...)
			}
			tasksCount--
		}

		var reports Reports
		reports = goReports

		Log(fmt.Sprintf("DONE: %v", reports.RenderJSON()))

		// - Download Files
		//  - Index code file with (Lines, Language, ProjectId, Filename)
		//  - Analyse code for smells and add to notification messages.
		//  - Analyse typical metrics and add to stats
		// - Find showcase file and save tabs, stats and metrics to redis

		ReportDone(r)
	}
}
