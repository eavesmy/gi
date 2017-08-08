package source

import (
	"fmt"
	"github.com/opesun/goquery"
)

func (o *One) ParseHTML(url string) bool {
	doc, err := goquery.ParseUrl(url)

	if err != nil {
		return false
	}

	go parseURL(doc)

	return true
}

func parseURL(doc goquery.Nodes) {

	hrefs := doc.Find("a")

	for i, _ := range hrefs {

		url := hrefs.Eq(i).Attr("href")

		fmt.Println(url)

	}

}
