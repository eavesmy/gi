package source

import (
	"../manager"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	// "github.com/zemirco/couchdb"
	"strings"
)

type One struct {
	Doc  *goquery.Document
	Urls int
	Url  string
}

func (o *One) getHtml(url string) (*goquery.Document, error) {

	resp, err := manager.NewClient(url)
	doc, _ := goquery.NewDocumentFromResponse(resp)

	if resp.StatusCode != 200 {
		doc = nil
	}

	return doc, err
}

func (o *One) urlSave(url string) {

	url = o.formatUrl(url)

	code := DB_Redis.SAdd("REDIS_URL_SAVE", url).Val()

	if code == 1 {
		o.Urls++
	}
}

func (o *One) parseHtml(tags []string) {

	for _, k := range tags {

		o.Doc.Find(k).Each(func(i int, d *goquery.Selection) {
			fmt.Println(d, k, "\n")
		})
	}

	o.Doc.Find("a").Each(func(i int, d *goquery.Selection) {
		href, _ := d.Attr("href")

		o.urlSave(href)
	})
}

func (o *One) done() {

	DB_Redis.SAdd("REDIS_URL_DONE", o.Url)

	url := o.getUrl()

	if url != "" {
		UrlList <- url
	}
}

func (o *One) getUrl() string {
	urls := DB_Redis.SDiff("REDIS_URL_SAVE", "REDIS_URL_DONE").Val()

	if len(urls) == 0 {
		return ""
	}

	return urls[0]
}

func (o *One) formatUrl(url string) string {

	if !strings.Contains(url, "http") && !strings.Contains(url, "www") {
		return o.Url + url
	}

	if !strings.Contains(url, "http") && strings.Contains(url, "www") {
		return "http:" + url
	}

	return url
}

func RunOne(url string, tags []string) {

	fmt.Println("RUN THIS ->", url)

	one := &One{}

	doc, err := one.getHtml(url)

	if doc == nil || err != nil {

		DoingTask.Faild++

		return
	}

	one.Doc = doc
	one.Urls = 0
	one.Url = url

	one.parseHtml(tags)

	one.done()

	DoingTask.Compelete++

	fmt.Println(DoingTask)
	return
}
