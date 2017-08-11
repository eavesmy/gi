package source

import (
	// "fmt"
	"github.com/opesun/goquery"
	"sync"
)

type ParseBody struct {
	Url  string
	Body map[string]string
}

func (o *One) ParseHTML(url string) *ParseBody {

	// fmt.Println("GET THIS URL ->", url)
	doc, _ := goquery.ParseUrl(url)

	dataBody := &ParseBody{}
	dataBody.Url = url

	var todo sync.WaitGroup
	todo.Add(2)

	go parseURL(doc, todo)
	go parseINFO(doc, todo, dataBody, o.Keys)

	todo.Wait()

	return dataBody
}

func parseURL(doc goquery.Nodes, todo sync.WaitGroup) {

	defer todo.Done()

	hrefs := doc.Find("a")

	for i, _ := range hrefs {

		url := hrefs.Eq(i).Attr("href")

		SaveURL(FormatURL(url))

	}

}

func parseINFO(doc goquery.Nodes, todo sync.WaitGroup, dataBody *ParseBody, keys *[]string) {

	defer todo.Done()

}
