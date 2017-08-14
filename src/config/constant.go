package config

var m = map[string]string{
	"STATUS_OK":         "200",
	"MAX_CLIENT":        "50",
	"REDIS_HOST":        "localhost:8004",
	"REDIS_URL_DB":      "1",
	"REDIS_URL_DONE_DB": "2",
	"SERVER_PORT":       ":8005",
	"URL_DB":            "1",
	"TASK_DONE":         "1",
	"TASK_DOING":        "2",
}

func Get(k string) string {
	return m[k]
}
