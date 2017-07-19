package source

import (
	// "fmt"
	"github.com/zdy23216340/gtool"
)

var queue []string
var dealingCount = 0

func Start(urls []string, tags []string) {
	for _, url := range urls {
		InsertURL(url)
	}
	getURL()

	for _, url := range queue {
		res := Http(url)

		status := DealRes(res)
		status.Url = url

		UpdateURL(status)
	}
}

func getURL() {
	urls := GetURL()

	for _, url := range urls {

		if gtool.GetIndex(queue, url) > -1 {
			continue
		}

		queue = append(queue, url)
	}
}
