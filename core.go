package gi

import (
	"net/http"
)

type Gi struct {
	Config *Config
	Client *http.Client

	URL_Root string

	URL_Total  int32
	URL_Failed int32
	URL_Done   int32

	URL_Filter []string // 过滤url中包含制定字符的
	URL_Only   []string

	Cookies []string
}

func New(rootUrl string, conf ...Config) *Gi {

	g := &Gi{
		URL_Root:   rootUrl,
		config:     &Config{},
		client:     &http.Client{},
		URL_Total:  0,
		URL_Failed: 0,
		URL_Done:   0,
		URL_Filter: []string{},
		URL_Only:   []string{},
		Cookies:    []string{},
	}

	if len(conf) > 0 {
		g.config = conf[0]
	}

	g.config.SetDefault()

	return g
}

func (g *Gi) Filter(str string) *Gi {
	g.URL_Filter = append(g.URL_Filter, str)
	return g
}

func (g *Gi) Only(str string) *Gi {
	g.URL_Only = append(g.URL_Only, str)
	return g
}

func (g *Gi) Go() {
}

func (g *Gi) Info() {

}

func (g *Gi) Handler() {

}
