package source

import (
	"strings"
)

func FormatUrl(url string) string {
	if !strings.Contains("http") && !strings.Contains("www") {
		return "http://"
	}
}
