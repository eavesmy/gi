package config

var m = map[string]string{
	"STATUS_OK":          "200",
	"MAX_CLIENT":         "50",
	"REDIS_HOST":         "localhost:8004",
	"REDIS_URL_DB":       "1",
	"REDIS_URL_CACHE_DB": "2",
	"SERVER_PORT":        ":8005",
	"URL_UNREPTIL":       "1",
	"URL_REPTILED":       "2",
	"URL_DONE":           "3",
}

func Get(k string) string {
	return m[k]
}
