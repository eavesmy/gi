package manager

import (
	"../config"
	"encoding/json"
	"github.com/go-redis/redis"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Info struct {
	Domin string `json:domin`
	Main  string `json:main`
}

func GetBody(req *http.Request) *Info {

	body, _ := ioutil.ReadAll(req.Body)

	req.Body.Close()

	info := &Info{}

	json.Unmarshal([]byte(body), info)

	return info
}

func NewRedisClient(db string) *redis.Client {

	dbNum, _ := strconv.Atoi(db)

	return redis.NewClient(&redis.Options{
		Addr:     config.Get("REDIS_HOST"),
		Password: "",
		DB:       dbNum,
	})
}
