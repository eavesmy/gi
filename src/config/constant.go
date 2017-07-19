package config

var m = map[string]string{
	"STATUS_OK":    "200",
	"MAX_CLIENT":   "50",
	"REDIS_HOST":   "localhost:8004",
	"REDIS_URL_DB": "1",
}

func Get(k string) string {
	return m[k]
}
