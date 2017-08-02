package manager

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
