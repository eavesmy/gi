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

	http.HandleFunc("/", source.Start)

	http.ListenAndServe(config.Get("SERVER_PORT"), nil)
}

func getConfig() ([]string, *[]string) {

	var Urls []string
	KeyWords := make([]string, 0)
	keys := &KeyWords

	fi, err := os.Open(CONFIG_PATH)

	if err != nil {
		panic(err)
	}

	defer fi.Close()

	buf := bufio.NewReader(fi)

	for {
		_buf, _, state := buf.ReadLine()

		if state == io.EOF {
			break
		}

		str := string(_buf)

		if strings.Contains(str, "http") {

			Urls = append(Urls, str)

		} else {

			*keys = append(*keys, str)
		}
	}

	return Urls, keys
}
