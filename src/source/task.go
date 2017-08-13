package source

import (
	"../manager"
	// "fmt"
	"strings"
)

type One struct {
	State    string
	Id       string
	Domin    string
	Urls     int
	Complete int
	Status   bool
	Undo     *[]string
	Done     int
	Keys     *[]string
}

type SaveData struct {
	Url  string
	Info map[string]string
}

var UrlList = make(chan string, MaxDealProcess)
var InfoList = make(chan SaveData, 1)
var DoingTask *One

func (o *One) AddUrl_() {
	o.Urls++
}

func (o *One) Complete_(url string) {
	DoneURL(url)
	o.Complete++
}

func (o *One) Run_(url string) {

	o.AddUrl_()

	url = FormatURL(url)

	o.ParseHTML(url)

	o.Complete_(url)
}

func NewTask(body *manager.Info) {

	one := &One{}

	keys := make([]string, 0)
	one.Keys = &keys

	DoingTask = one

	one.Domin = FormatDomin(body.Domin)

	for _, k := range strings.Split(body.Main, ",") {
		*one.Keys = append(*one.Keys, k)
	}

	SaveURL(one.Domin)

	GetURL()

	for {
		go one.Run_(<-UrlList)
	}

}
