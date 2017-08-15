package source

import (
	"../manager"
	"strings"
)

const maxDealProcess = 5

var DoingTask *Task

type Task struct {
	Total     int
	Compelete int
	Faild     int
	Domin     string
	Tags      []string
	State     string
}

// var TaskList = make(chan *One, maxDealProcess)
var UrlList = make(chan string, maxDealProcess)

func (t *Task) runTask() {

	for {
		go RunOne(<-UrlList, t.Tags)
	}
}

func (t *Task) GetState() {
}

func _task() {

}

func NewTask(body *manager.Info) bool {

	if DoingTask != nil {
		return false
	}

	task := &Task{}

	task.Faild = 0
	task.Compelete = 0
	task.Total = 0

	task.Domin = body.Domin
	task.Tags = func() []string {
		return strings.Split(body.Main, ",")
	}()

	DoingTask = task

	UrlList <- task.Domin

	go task.runTask()

	return true
}
