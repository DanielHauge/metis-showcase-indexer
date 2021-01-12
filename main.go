package main

import (
	. "./coordinator"
	. "./shared"
	. "./worker"
	"errors"
	"os"
)

func main() {
	mode := os.Getenv("mode")
	switch mode {
		case Worker: RunWorker()
		case Coordinator: RunCoordinator()
		default: panic(errors.New("No default mode found, please set environment variable: 'mode' like -e mode=worker"))
	}
}