package source

import (
	"../manager"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type One struct {
	Doc  *goquery.Document
	Urls int
	Url  string
}

func (o *One) getHtml(url string) (*goquery.Document, error) {
	return goquery.NewDocument(url)
}

func (o *One) urlSave(url string) {

	url = manager.FormatUrl(url)

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

	// fmt.Println(line, k)
}

func (o *One) done() {

	url := o.getUrl()

	if url != "" {
		UrlList <- url
	}
}

func (o *One) getUrl() string {
	// From save but not in done.
	// Do not remove from save.

	return ""
}

func RunOne(url string, tags []string) {

	one := &One{}

	doc, err := one.getHtml(url)

	if err != nil {

		DoingTask.Faild++

		return
	}

	one.Doc = doc
	one.Urls = 0
	one.Url = url

	one.parseHtml(tags)

	one.done()

	DoingTask.Compelete++

	return
}
