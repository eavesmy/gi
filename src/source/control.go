package source

import (
	"fmt"
	"net/http"
	"strings"
	// "github.com/zdy23216340/gtool"
)

var queue []string
var dealingCount = 0
var Info = map[string]string{}

func Start(w http.ResponseWriter, req *http.Request) {

	if !strings.Contains(req.RemoteAddr, "127.0.0.1") {
		return
	}

	config := req.ParseForm()

	fmt.Println(config)
}

func parse(keys *[]string) {

	queue = GetURL(5)

	if len(queue) == 0 {
		fmt.Println("All HTML parse done.")
		return
	}

	for i := 0; i < len(queue); i++ {
		url := queue[i]

		res := Http(url)

		fmt.Println(res.Header["meta"])
		// status := DealRes(res, keys)
		// status.Url = url

		// UpdateURL(status)
	}

	parse(keys)
}

func parseInfoStruct(keys *[]string) {

	for _, key := range *keys {

		Info[key] = ""
	}
}
