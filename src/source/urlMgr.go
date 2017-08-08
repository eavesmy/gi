package source

import (
	// "fmt"
	"strings"
)

func FormatURL(url string) string {

	if !strings.Contains(url, "http") {
		url = "http://" + url
	}

	return url
}
