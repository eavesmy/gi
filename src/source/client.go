package source

import (
	"net/http"
)

func Http(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "User-Agent:Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")

	client := &http.Client{}

	client.Get(url)

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	return res
}
