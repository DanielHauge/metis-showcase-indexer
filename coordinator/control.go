package coordinator

import (
	. "../shared"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


func ControlServer(){
	var workerId string
	var status string
	var since string
	var s string

	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		b, err := ioutil.ReadAll(request.Body)
		if err != nil { http.Error(writer, err.Error(), 500); return }
		defer request.Body.Close()
		s := string(b)
		repositoriesSpace.Put(s, time.Now().Add(1*time.Second).Format(TimeFormat))
		Log("Repository added")
	})

	http.HandleFunc("/log", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		Log("Logs requested")
		query, _ := LogSpace.QueryAll(&s)
		for _, t := range query {
			writer.Write([]byte(t.GetFieldAt(0).(string)))
		}
	})

	http.HandleFunc("/tasks", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		Log("Logs requested")
		query, _ := IndexSpace.QueryAll(&s)
		for _, t := range query {
			writer.Write([]byte(t.GetFieldAt(0).(string)+"\n"))
		}
	})

	http.HandleFunc("/repo", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		Log("Logs requested")

		query, _ := repositoriesSpace.QueryAll(&s, &s)
		for _, t := range query {
			writer.Write([]byte(t.GetFieldAt(0).(string)+ " " +t.GetFieldAt(1).(string)+"\n"))
		}

	})

	http.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		Log("Status requested")
		query, _ := StatusSpace.QueryAll(&workerId, &status, &since)
		for _, t := range query {
			writer.Write([]byte(t.String()+"\n"))
		}
	})

	Log("Control server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
