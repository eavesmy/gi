package source

import (
	"../config"
	// "fmt"
	"github.com/go-redis/redis"
	"strconv"
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

func InsertURL(url string) {

	err := redisClient().Set(url, "no", 0).Err()

	if err != nil {
		panic(err)
	}
}

func GetURL() []string {

	var notFinishedUrls []string

	urls := redisClient().Keys("*").Val()

	for _, url := range urls {
		isDone := redisClient().Get(url).Val()

		if isDone == "yes" {
			continue
		}

		notFinishedUrls = append(notFinishedUrls, url)

		if len(notFinishedUrls) > 5 {
			break
		}
	}

	return notFinishedUrls
}

func UpdateURL(status *Status) {
	//更新url
	if status.GetHtmlCode {
		redisClient().Set(status.Url, "yes", 0)
	} else {
		redisClient().Set(status.Url, "faild", 0)
	}
}
