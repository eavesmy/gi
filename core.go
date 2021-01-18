package gi

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Gi struct {
	Config *Config
	Client *Client

	URL_Root string
	URL_Host string

	URL_Filter []string // 过滤url中包含制定字符的
	URL_Only   []string

	url_Chan chan string

	Cookies []*http.Cookie
	Headers map[string]string

	Cache *Cache

	Handlers []func(*Context)
}

var chan_exit = make(chan bool, 1)

func New(conf ...*Config) *Gi {

	g := &Gi{
		Config:     &Config{Domain: true},
		Client:     NewClient(),
		URL_Filter: []string{},
		URL_Only:   []string{},
		Cookies:    []*http.Cookie{},
		Cache:      NewCache(),
		Headers:    map[string]string{},
		url_Chan:   make(chan string, 1),
	}

	if len(conf) > 0 {
		g.Config = conf[0]
	}

	g.Config.SetDefault()

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

func (g *Gi) Go(_url string) {

	if _url == "" {
		fmt.Println("Param url must be exists")
		return
	}

	u, err := url.Parse(_url)
	if err != nil {
		fmt.Println(err)
		return
	}

	g.URL_Host = u.Host

	g.URL_Root = _url

	g.Cache._Chan = g.url_Chan
	go g.Cache.Go()

	g.Cache.Add(_url)

	for i := 0; i < g.Config.MaxRuntime; i++ {
		go g.run()
	}

	// g.Cache.InfoLoop(10)
	g.stay()
}

func (g *Gi) Info() {

}

func (g *Gi) Handler(handler func(*Context)) {
	g.Handlers = append(g.Handlers, handler)
}

func (g *Gi) stay() {
	for {
		<-chan_exit
	}
}

func (g *Gi) run() {
	for {
		url := <-g.url_Chan

		g.mainProgram(url)
	}
}

func (g *Gi) mainProgram(_url string) {

	g.Cache.Doing(_url)

	req, err := g.Client.Do(g.Config.Method, _url)

	if err != nil {
		g.Cache.Failed(_url, err)
		return
	}

	for k, v := range g.Headers {
		req.Header.Add(k, v)
	}

	for _, cookie := range g.Cookies {
		req.AddCookie(cookie)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	res, err := (&http.Client{Transport: tr}).Do(req)

	if err != nil {
		fmt.Println("request err: ", err, _url)
		g.Cache.Failed(_url, err)
		return
	}

	ctx := &Context{Res: res, Req: req, Gi: g}

	nodes, err := ctx.Html()

	ctx.Nodes = nodes

	// 获取页面上所有链接

	if g.Config.AutoParseUrl {
		hrefs := ctx.GetHref()

		for _, href := range hrefs {

			// 判断域
			if isDomain, _ := g.url_domain(href); g.Config.Domain && !isDomain {
				continue
			}

			if href == "javascript:" || href == _url+"/" || href == "/" {
				continue
			}

			if strings.IndexAny(href, "#") == 0 && !strings.ContainsAny(href, "/") {
				continue
			}

			// 判断链接合法

			if g.url_only(href) && !g.url_filter(href) {

				u, _ := url.Parse(href)
				if u.Host == "" {
					href = g.URL_Root + href
				}
				g.Cache.Add(href)
			}
		}
	}

	for _, handler := range g.Handlers {
		handler(ctx)
	}

	g.Cache.Done(_url)
}

func (g *Gi) url_domain(_url string) (bool, error) {

	u, err := url.Parse(_url)

	// fmt.Println(u.Host, g.URL_Host, "url", _url)

	if err != nil {
		return false, err
	}

	if u.Host == "" || u.Host == g.URL_Host {
		return true, nil
	}

	return false, nil
}

func (g *Gi) url_only(url string) bool {
	for _, str := range g.URL_Only {
		if strings.Contains(url, str) {
			// fmt.Println("is true", url, str)
			return true
		}
	}
	return false
}

func (g *Gi) url_filter(url string) bool {
	for _, str := range g.URL_Filter {
		if strings.Contains(url, str) {
			return true
		}
	}
	return false
}

func (g *Gi) AddCookie(str interface{}) {

	switch str.(type) {
	case http.Cookie:
		g.Cookies = append(g.Cookies, str.(*http.Cookie))
	case string:
		// name,value,httponly,maxAge,
		arr_str := strings.Split(str.(string), " ")
		for _, cookie := range arr_str {
			arr_c := strings.Split(cookie, ",")
			if len(arr_c) < 5 {
				fmt.Println("You can add symple cookie here but need \"name,value,httpOnly,maxAge\"")
				return
			}

			_bool, _ := strconv.ParseBool(arr_c[2])
			_int, _ := strconv.Atoi(arr_c[3])

			c := &http.Cookie{
				Name:     arr_c[0],
				Value:    arr_c[1],
				HttpOnly: _bool,
				MaxAge:   int(_int),
			}

			g.Cookies = append(g.Cookies, c)
		}
	default:
		fmt.Println("param type must be http.Cookie or string")
		return
	}
}

func (g *Gi) AddHeader(k string, v string) {
	g.Headers[k] = v
}
