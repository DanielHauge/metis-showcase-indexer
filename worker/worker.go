package worker

import (
	. "../shared"
	"fmt"
	"github.com/beevik/guid"
)

var workerId string

func RunWorker() {
	workerId = guid.New().String()
	ClientName = fmt.Sprintf("Worker: %v", workerId)
	ConnectLogSpace()
	ConnectWorkerStatus()
	defer CloseStatusSpace()
	InitTaskReceiver()

}
