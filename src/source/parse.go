package source

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type One struct {
	Doc *goquery.Document
}

func (o *One) getHtml(url string) (*goquery.Document, error) {
	return goquery.NewDocument(url)

}

func (o *One) parseHtml(tags []string) {

	for _, k := range tags {

		line := o.Doc.Find(k).Text()

	}

}

func (o *One) saveUrl() {

}

func ParseNext() {

}

func RunOne(url string, tags []string) {

	one := &One{}

	doc, err := one.getHtml(url)

	if err != nil {

		DoingTask.Faild++

		return
	}

	one.Doc = doc

	one.parseHtml(tags)

	// one.saveUrl()

	DoingTask.Compelete++
}
