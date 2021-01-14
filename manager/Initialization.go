package manager

import (
	. "../shared"
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing/object"
	"github.com/go-git/go-git/storage/memory"
	"path/filepath"
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

		TaskSpace.Put(ClientName, TASK, CLEAR, r, "none", "none")

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
			fileExt := filepath.Ext(n)

			if !CheckError(e){
				TaskSpace.Put(ClientName, TASK, INDEX_FILE, r, n, content )
				if fileExt == ".go" {
					TaskSpace.Put(ClientName, TASK, ANALYSE_GO, r, n, content )
				}
				if fileExt == ".showcase" {
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
		TaskSpace.Put(ClientName, TASK, INDEX_REPORT, r, "report.analysis", reports.RenderJSON())


		ReportDone(r)
	}
}
