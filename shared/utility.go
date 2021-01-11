package main

import "os"

func NetworkArgs() (string,string,string) {
	host := withDefault(os.Getenv("host"), "localhost")
	port := withDefault(os.Getenv("port"), "31415")
	space := withDefault(os.Getenv("space"), "distributed")
	return host, port, space
}

func SpaceUri() string{
	host, port, space := NetworkArgs()
	return "tcp://" + host + ":" + port + "/" + space
}

func withDefault(s string, def string) string{
	if len(s) == 0 { return def }
	return s
}
