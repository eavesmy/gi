package source

import (
	"github.com/opesun/goquery"
	"sync"
)

type ParseBody struct {
	Url  string
	Body map[string]string
}

func (o *One) ParseHTML(url string) *ParseBody {

	doc, _ := goquery.ParseUrl(url)

	dataBody := &ParseBody{}

	var todo sync.WaitGroup
	todo.Add(2)

	go parseURL(doc, todo, dataBody)
	go parseINFO(doc, todo, dataBody)

	todo.Wait()

	return dataBody
}

func parseURL(doc goquery.Nodes, todo sync.WaitGroup, dataBody *ParseBody) {

	defer todo.Done()

	hrefs := doc.Find("a")

	for i, _ := range hrefs {

		url := hrefs.Eq(i).Attr("href")

		SaveURL(FormatURL(url))

	}

}

func parseINFO(doc goquery.Nodes, todo sync.WaitGroup, dataBody *ParseBody) {
	defer todo.Done()
}
