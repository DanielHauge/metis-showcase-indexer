package shared

import (
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"os"
)



func LogSpaceUri() string{
	return "tcp://" + withDefault(os.Getenv("host"), "localhost") + ":9093/log"
}

func TaskSpaceUri() string{
	return "tcp://" + withDefault(os.Getenv("host"), "localhost") + ":9094/task"
}

func IndexSpaceUri() string {
	return "tcp://" + withDefault(os.Getenv("host"), "localhost") + ":9091/index"
}


func ConnectTaskSpace(){
	uri := TaskSpaceUri()
	TaskSpace = space.NewRemoteSpace(uri)
	Log(fmt.Sprintf("Connected to task space at: %v", uri))
}

func ConnectLogSpace(){
	uri := LogSpaceUri()
	LogSpace = space.NewRemoteSpace(uri)
	Log(fmt.Sprintf("Connected to log space at: %v", uri))
}



