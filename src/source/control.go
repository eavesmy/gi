package source

import (
	"../manager"
	"net/http"
	"net/url"
	// "github.com/zdy23216340/gtool"
)

var queue []string
var dealingCount = 0
var Info = map[string]string{}

func Start(w http.ResponseWriter, req *http.Request) {

	if !fromLocal(req.Host) {
		return
	}

	body := manager.GetBody(req)

	InsertURL(body.Domin)
}

func fromLocal(host string) bool {
	_url, _ := url.Parse("http://" + host)

	_host := _url.Hostname()

	return (_host == "localhost" || _host == "127.0.0.1")
}
