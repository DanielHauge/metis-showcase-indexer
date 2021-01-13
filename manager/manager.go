package manager

import (
	. "../shared"
	"fmt"
	"github.com/beevik/guid"
)

var managerId string

func RunManager() {
	managerId = guid.New().String()
	ClientName = fmt.Sprintf("Manager: %v", managerId)
	ConnectLogSpace()
	defer ConnectStatusSpace(managerId)()
	InitManager()
}


