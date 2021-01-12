package worker

import (
	. "../shared"
	"github.com/rs/xid"
)

var workerId string

func RunWorker() {

	workerId = xid.New().String()
	Logf("Initilizing worker with id: %v",workerId)
	ConnectWorkerStatus()
	defer CloseStatusSpace()
	InitTaskReceiver()

}
