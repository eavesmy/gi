package main

import (
	"./config"
	"./source"
	"net/http"
)

func main() {
	http.HandleFunc("/api/reptile/go", source.Start)

	http.ListenAndServe(config.Get("SERVER_PORT"), nil)
}
