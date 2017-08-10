package config

var m = map[string]string{
	"STATUS_OK":    "200",
	"MAX_CLIENT":   "50",
	"REDIS_HOST":   "localhost:8004",
	"REDIS_URL_DB": "1",
	"SERVER_PORT":  ":8005",
	"URL_UNEXISTS": "2",
}

func Get(k string) string {
	return m[k]
}
