package gi

import (
	"bytes"
	"net/http"
)

// 实现目标：
// 对 root url 进行爬取，限定爬度url在root内
// 设定url过滤条件

type Client struct {
	c          *http.Client
	RunTimeNum int
}

func NewClient() *Client {
	return &Client{
		c: &http.Client{},
	}
}

func (c *Client) Do(method string, url string, datas ...[]byte) (*http.Request, error) {
	var data []byte
	if len(datas) > 0 {
		data = datas[0]
	} else {
		data = []byte{}
	}

	postData := bytes.NewBuffer(data)

	return http.NewRequest(method, url, postData)
}
