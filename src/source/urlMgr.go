package source

import (
	"../config"
	// "fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
)

func redisClient() *redis.Client {

	db, _ := strconv.Atoi(config.Get("REDIS_URL_DB"))

	client := redis.NewClient(&redis.Options{
		Addr:     config.Get("REDIS_HOST"),
		Password: "", // no password set
		DB:       db,
	})

	return client
}

func InitURL(url string) {

}

func InsertURL(url string) {

	client := redisClient()
	isDone := client.Get(url).Val()

	defer client.Close()

	if isDone != "" {
		return
	}

	err := client.Set(url, "no", 0).Err()

	if err != nil {
		panic(err)
	}
}

func GetURL(count int) []string {

	var notFinishedUrls []string

	client := redisClient()
	urls := client.Keys("*").Val()

	defer client.Close()

	for _, url := range urls {
		isDone := client.Get(url).Val()

		if isDone == "yes" {
			continue
		}

		if !strings.Contains(url, "http") {
			url = "http://" + url
		}

		notFinishedUrls = append(notFinishedUrls, url)

		if len(notFinishedUrls) > count {
			break
		}
	}

	return notFinishedUrls
}

func UpdateURL(status *Status) {

	client := redisClient()

	if status.GetHtmlCode {
		client.Set(status.Url, "yes", 0)
	} else {
		client.Set(status.Url, "faild", 0)
	}
	defer client.Close()
}
