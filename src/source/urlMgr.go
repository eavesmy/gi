package source

import (
	"../config"
	"fmt"
	"strings"
)

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

	v, err := NewRedisClient.Get(url).Result()

	if err != nil {

		_err := NewRedisClient.Set(url, config.Get("URL_UNEXISTS"), 0).Err()

		if _err != nil {
			fmt.Println("Insert url failed")
		}
		return false
	}

	fmt.Println(111, v)

	return true
}
