package source

import (
	"../manager"
)

const maxDealProcess = 5

var programState *Task

type Task struct {
	Total     int
	Compelete int
	Faild     int
	Domin     string
	Keys      string
	State     int
}

// var TaskList = make(chan *One, maxDealProcess)
var UrlList = make(chan string, maxDealProcess)

func (t *Task) runTask() {

	for {
		url := <-UrlList
	}
}

func _task() {

}

func NewTask(body *manager.Info) bool {

	if programState != nil {
		return false
	}

	task := &Task{}

	task.Domin = body.Domin
	task.Keys = body.main

	UrlList <- task.Domin

	go task.runTask()

	return true
}
