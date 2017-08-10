package source

import (
	"../manager"
	"github.com/go-redis/redis"
	"net/http"
	"net/url"
	// "github.com/zdy23216340/gtool"
)

const MaxDealProcess = 5

var DealingCount = 0
var NewRedisClient *redis.Client

func Start(w http.ResponseWriter, req *http.Request) {

	if !fromLocal(req.Host) {
		return
	}

	body := manager.GetBody(req)

	if DoingTask != nil {
		w.Write([]byte("Already has a task running"))
		return
	}

	// Init redis & couchdb
	NewRedisClient = manager.NewRedisClient()

	go NewTask(body)

	w.Write([]byte("Task Added"))
}

func fromLocal(host string) bool {
	_url, _ := url.Parse("http://" + host)

	_host := _url.Hostname()

	return (_host == "localhost" || _host == "127.0.0.1")
}
