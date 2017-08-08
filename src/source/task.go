package source

import (
	"../manager"
	"strings"
)

var UrlList = make(chan string, 5)

type One struct {
	State    string
	Id       string
	Domin    string
	Urls     int
	Complete int
	Status   bool
	Undo     []string
	Done     []int
	Keys     []string
}

func (o *One) AddUrl_() {
	o.Urls++
}

func (o *One) Complete_() {
	o.Complete++
}

func (o *One) Save_() {

}

func (o *One) Run_(url string) {

	o.AddUrl_()

	url = FormatURL(url)

	if o.ParseHTML(url) {

		o.Complete_()

	}

}

func NewTask(body *manager.Info) {
	one := &One{}
	one.Domin = body.Domin

	for _, k := range strings.Split(body.Main, ",") {
		one.Keys = append(one.Keys, k)
	}

	UrlList <- one.Domin

	for {
		go one.Run_(<-UrlList)
	}
}

// Need a list to put url. if list full then wait to put in.
//	Insert url
//	Check the process num.
//	if process num < limit then

// Every time update ->
// 1. Url num.
// 2.
