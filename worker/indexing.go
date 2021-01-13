package worker

import (
	. "../shared"
	"fmt"
)

func IndexFile(repo string, file string, content string){
		Log(fmt.Sprintf("Indexing %v", file))
}
