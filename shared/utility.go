package shared

import (
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"os"
	"time"
)

const Worker = "worker"
const Coordinator = "coordinator"
const TimeFormat = "2006-01-02 15:04:05"
var LogSpace space.Space
var ClientName string

type StatusReport struct {
	WorkerId string
	Status string
	Since time.Time
}

func StatusSpaceUri() string{
	return "tcp://" + withDefault(os.Getenv("host"), "localhost") + ":9092/status"
}

func LogSpaceUri() string{
	return "tcp://" + withDefault(os.Getenv("host"), "localhost") + ":9093/log"
}

func TaskSpaceUri() string {
	return "tcp://" + withDefault(os.Getenv("host"), "localhost") + ":9091/task"
}

func withDefault(s string, def string) string{
	if len(s) == 0 { return def }
	return s
}

func Log(message string){
	log := fmt.Sprintf("%v - %v\n", ClientName, message)
	fmt.Println(log)
	LogSpace.Put(log)
}


