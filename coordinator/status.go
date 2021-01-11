package coordinator

import (
	"fmt"
	"github.com/DanielHauge/goSpace/space"
	"time"
)

var statusSpace space.Space

type StatusReport struct {
	WorkerId string
	Status string
	Since time.Time
}

func InitStatusSpace(uri string){
	statusSpace = space.NewSpace(uri)
}

func GetStatusReports(){
	workerId := ""
	status := ""
	since := time.Now()

	query, _ := statusSpace.QueryAll(&workerId, &status, &since)
	for _, t := range query {
		fmt.Println(t.String())
	}
}

