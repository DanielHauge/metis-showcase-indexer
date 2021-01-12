package coordinator

import (
	. "../shared"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


func ControlServer(){

	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		b, err := ioutil.ReadAll(request.Body)
		if err != nil { http.Error(writer, err.Error(), 500); return }
		defer request.Body.Close()
		s := string(b)
		repositoriesSpace.Put(s, time.Now().Add(10*time.Second).Format(TimeFormat))
		Log("Repository added")
	})

	http.HandleFunc("/log", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		Log("Logs requested")
		GetLogs(writer)
	})

	http.HandleFunc("/tasks", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		Log("Logs requested")
		GetTasks(writer)
	})

	http.HandleFunc("/repo", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		Log("Logs requested")
		GetRepositories(writer)
	})

	http.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		Log("Status requested")
		GetStatusReports(writer)
	})

	Log("Control server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
