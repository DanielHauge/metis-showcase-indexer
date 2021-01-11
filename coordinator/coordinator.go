package coordinator

import . "../shared"

var Repositories = make(map[string]string)

func Coordinator() {

	go ControlServer()
	InitStatusSpace(StatusSpace())


	<- make(chan string)
}


