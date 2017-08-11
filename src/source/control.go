package source

import (
	"../config"
	"../manager"
	"github.com/go-redis/redis"
	"net/http"
	"net/url"
	// "github.com/zdy23216340/gtool"
)

const MaxDealProcess = 5

var DealingCount = 0
var UrlClient *redis.Client
var UrlDoneClient *redis.Client

func Start(w http.ResponseWriter, req *http.Request) {

	if !fromLocal(req.Host) {
		return
	}

	body := manager.GetBody(req)

	if DoingTask != nil {
		w.Write([]byte("Already has a task running"))
		return
	}

	//Init redis client
	UrlClient = manager.NewRedisClient(config.Get("REDIS_URL_DB"))
	UrlDoneClient = manager.NewRedisClient(config.Get("REDIS_URL_DONE_DB"))

	go NewTask(body)

	w.Write([]byte("Task Added"))
}

func fromLocal(host string) bool {
	_url, _ := url.Parse("http://" + host)

	_host := _url.Hostname()

	return (_host == "localhost" || _host == "127.0.0.1")
}
