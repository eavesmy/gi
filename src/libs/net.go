package libs

import (
	"net/http"
)

func NewHttpClient() *http.Client {
	return &http.Client{}
}
