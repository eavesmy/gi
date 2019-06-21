package gi

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

type Context struct {
	Res   *http.Response
	Req   *http.Request
	Nodes *goquery.Document
}

func (c *Context) String() (string, error) {
	bytes, err := ioutil.ReadAll(c.Res.Body)
	return string(bytes), err
}

func (c *Context) Bytes() ([]byte, error) {
	return ioutil.ReadAll(c.Res.Body)
}

func (c *Context) Html() (*goquery.Document, error) {

	// Parse by github.com/PuerkitoBio/goquery

	doc, err := goquery.NewDocumentFromReader(c.Res.Body)

	if err != nil {
		fmt.Println(err, doc)
		return nil, err
	}

	return doc, nil
}

func (c *Context) GetPath() string {
	return c.Req.URL.Path
}

func (c *Context) GetHref() []string {

	_map := []string{}

	c.Nodes.Find("a").Each(func(i int, a *goquery.Selection) {
		attr, exists := a.Attr("href")
		if exists {
			_map = append(_map, attr)
		}
	})

	return _map
}
