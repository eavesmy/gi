package manager

import (
	"../config"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"net/http"
	"net/url"
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

func NewClient(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Check err ->", err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36")

	return client.Do(req)
}

func NewRedisClient(db string) *redis.Client {

	dbNum, _ := strconv.Atoi(db)

	return redis.NewClient(&redis.Options{
		Addr:     config.Get("REDIS_HOST"),
		Password: "",
		DB:       dbNum,
	})
}

func FromLocal(host string) bool {
	_url, _ := url.Parse("http://" + host)

	_host := _url.Hostname()

	return (_host == "localhost" || _host == "127.0.0.1")

}
