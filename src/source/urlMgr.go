package source

import (
	"../config"
	"../manager"
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

	c := manager.NewRedisClient(config.Get("URL_DB"))

	defer c.Close()

	err := c.SAdd("REDIS_URL_PREPARE_LIST", url).Err()

	if err != nil {
		fmt.Println("Save url err ->", err)

		return false
	}

	return true
}

func GetURL() {
	url := getURL()

	if url != "" {
		UrlList <- url
	}
}

func getURL() string {
	c := manager.NewRedisClient(config.Get("URL_DB"))

	defer c.Close()

	return c.SPop("REDIS_URL_PREPARE_LIST").Val()

}

func DoneURL(url string) {

	c := manager.NewRedisClient(config.Get("URL_DB"))

	defer c.Close()

	state, err := c.SMove("REDIS_URL_PREPARE_LIST", "REDIS_URL_DONE_LIST", url).Result()

	if err != nil {
		fmt.Println("Done url err->", err)
	}

	fmt.Println("Done url ->", state)
}
