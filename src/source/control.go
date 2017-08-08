package source

import (
	"../manager"
	"net/http"
	"net/url"
	// "github.com/zdy23216340/gtool"
)

const MaxDealProcess = 5

var DealingCount = 0

func Start(w http.ResponseWriter, req *http.Request) {

	if !fromLocal(req.Host) {
		return
	}

	body := manager.GetBody(req)

	go NewTask(body)

	w.Write([]byte("oooooo"))
}

func fromLocal(host string) bool {
	_url, _ := url.Parse("http://" + host)

	_host := _url.Hostname()

	return (_host == "localhost" || _host == "127.0.0.1")
}
