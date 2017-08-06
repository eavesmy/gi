package main

import (
	"./config"
	"./source"
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
)

const CONFIG_PATH = "./config.txt"

func main() {

	http.HandleFunc("/api/reptile/go", source.Start)

	http.ListenAndServe(config.Get("SERVER_PORT"), nil)
}
