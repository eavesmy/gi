package source

import (
	"fmt"
	"strings"
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
	//TODO : Fix this.
	UrlClient.LPush("REDIS_URL_PREPARE_LIST", url)

	return false
}

func GetURL() string {
	url := GetURL()

	if url == "" {
		url = GetURL()
	}

	return url
}

func getURL() string {
	v := UrlClient.LIndex("REDIS_URL_PREPARE_LIST", -1).String()

	if v == UrlDoneClient.Get(v).String() {

		return ""

	} else {

		return v

	}
}

func DoneURL(url string) {
	err := UrlDoneClient.Set(url, true, 0).Err()

	if err != nil {
		fmt.Println("Done a url ->", url, ",but some thing error here.")
	}
}
