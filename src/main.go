package main

import (
	"./source"
	"bufio"
	"io"
	"os"
	"strings"
)

const CONFIG_PATH = "./config.txt"

func main() {

	urls, tags := getConfig()

	source.Start(urls, tags)
}

func getConfig() ([]string, []string) {

	var Urls []string
	var KeyWords []string

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

			KeyWords = append(KeyWords, str)
		}
	}

	return Urls, KeyWords
}
