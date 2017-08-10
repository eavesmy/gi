package source

import (
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

	v, err := UrlClient.Get(url).Result()
	//1.get 2.

	if err != nil && v == "" {

	}

	return false
}

//func GetURL() *[]string {
//}
