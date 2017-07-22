package source

import (
	"fmt"
	// "github.com/zdy23216340/gtool"
)

var queue []string
var dealingCount = 0

func Start(urls []string, keys []string) {

	for _, url := range urls {
		InsertURL(url)
	}

	parse(keys)
}

func parse(keys []string) {

	queue = GetURL(5)

	if len(queue) == 0 {
		fmt.Println("All HTML parse done.")
		return
	}

	for i := 0; i < len(queue); i++ {
		url := queue[i]

		fmt.Println("Deal url ->", url)
		res := Http(url)

		status := DealRes(res, keys)
		status.Url = url

		UpdateURL(status)
	}

	parse(keys)
}
