package source

import (
	"github.com/opesun/goquery"
	"sync"
)

func (o *One) ParseHTML(url string) bool {

	doc, err := goquery.ParseUrl(url)

	if err != nil {
		return false
	}

	var todo sync.WaitGroup
	todo.Add(2)

	go parseURL(doc, todo)
	go parseINFO(doc, todo)

	todo.Wait()

	return true
}

func parseURL(doc goquery.Nodes, todo sync.WaitGroup) {

	defer todo.Done()

	hrefs := doc.Find("a")

	for i, _ := range hrefs {

		url := hrefs.Eq(i).Attr("href")

		SaveURL(FormatURL(url))

	}

}

func parseINFO(doc goquery.Nodes, todo sync.WaitGroup) {
	defer todo.Done()

}
