package source

import (
	"fmt"
	"strings"
	"time"
)

type UrlState struct {
	State string
	Url   string
}

func FormatURL(url string) string {

	if !strings.Contains(url, "http") && !strings.Contains(url, "www") {
		url = DoingTask.Domin + url
	}

	if !strings.Contains(url, "http") && strings.Contains(url, "www") {
		url = "http:" + url
	}

	return url
}

func FormatDomin(url string) string {

	if !strings.Contains(url, "http") {
		url = "http://" + url
	}

	return url
}

func ParseDomin(url string) string {
	//TODO :  Use url api.

	return url
}

func SaveURL(url string) bool {

	UrlClient.LPush("REDIS_URL_PREPARE_LIST", url)

	return false
}

func GetURL() string {
	url := getURL()

	if url == "" {
		url = GetURL()
	}

	return url
}

func getURL() string {
	v := UrlClient.BLPop(0, "REDIS_URL_PREPARE_LIST").Val()

	if len(v) == 0 {
		return ""
	}

	if v[1] == UrlDoneClient.Get(v[1]).Val() {

		return ""

	} else {

		return v[1]

	}
}

func DoneURL(url string) {
	err := UrlDoneClient.Set(url, true, 0).Err()

	if err != nil {
		fmt.Println("Done a url ->", url, ",but some thing error here.")
	}
}
