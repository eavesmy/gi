package source

import (
	"../manager"
	"net/http"
	// "github.com/zdy23216340/gtool"
)

func Start(w http.ResponseWriter, req *http.Request) {

	if !manager.FromLocal(req.Host) {
		return
	}

	body := manager.GetBody(req)
	state := NewTask(body)

	if state {
		w.Write([]byte("Task Added"))
	} else {
		w.Write([]byte("Already has a task running"))
	}
}
