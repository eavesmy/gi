package source

import (
	"../config"
	"../manager"
	"github.com/go-redis/redis"
	"net/http"
	// "github.com/zdy23216340/gtool"
)

var DB_Redis *redis.Client

func Start(w http.ResponseWriter, req *http.Request) {

	if !manager.FromLocal(req.Host) {
		return
	}

	body := manager.GetBody(req)
	state := NewTask(body)

	// Init db
	DB_Redis = manager.NewRedisClient(config.Get("URL_DB"))

	if state {
		w.Write([]byte("Task Added"))
	} else {
		w.Write([]byte("Already has a task running"))
	}
}
