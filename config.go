package gi

type Config struct {
	Method string // GET POST
	Domain bool

	MaxRuntime int // 同时处理数量
	Timeout    int // 请求超时 unit:s
	RetryCount int // 重试次数

	AutoParseUrl bool
}

func (c *Config) SetDefault() {
	if c.Method == "" {
		c.Method = "GET"
	}
	if c.MaxRuntime == 0 {
		c.MaxRuntime = 1
	}
	if c.Timeout == 0 {
		c.Timeout = 5
	}
	if c.RetryCount == 0 {
		c.RetryCount = 3
	}
}
