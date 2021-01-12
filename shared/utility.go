package shared

import (
	"fmt"
	"os"
	"time"
)

const Worker = "worker"
const Coordinator = "coordinator"

type StatusReport struct {
	WorkerId string
	Status string
	Since time.Time
}

func networkArgs() (string,string) {
	host := withDefault(os.Getenv("host"), "localhost")
	port := withDefault(os.Getenv("port"), "31415")
	return host, port
}

func StatusSpace() string{
	host, port := networkArgs()
	return "tcp://" + host + ":" + port + "/status"
}


func TaskSpace() string {
	host, port := networkArgs()
	return "tcp://" + host + ":" + port + "/task"
}

func withDefault(s string, def string) string{
	if len(s) == 0 { return def }
	return s
}

func Log(s string){
	fmt.Println(s)
}

func Logf(s string, a ...interface{}){
	fmt.Printf(s, a)
}

