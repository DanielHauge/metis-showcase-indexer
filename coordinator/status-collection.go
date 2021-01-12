package coordinator

import (
	. "../shared"
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"time"
)

var statusSpace space.Space

func InitStatusSpace(uri string){
	statusSpace = space.NewSpace(uri)
	Log("Status space initialized")
}

func GetStatusReports(){
	var workerId string
	var status string
	var since time.Time

	query, _ := statusSpace.QueryAll(&workerId, &status, &since)
	for _, t := range query {
		fmt.Println(t.String())
	}
}

