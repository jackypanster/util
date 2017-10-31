package util

import (
	"log"
	"io/ioutil"
)

var JobQueue chan Job
var debug = false

func SetDebug(enable bool) {
	debug = enable
	if !debug {
		log.SetOutput(ioutil.Discard)
	}
}

func InitQueue(maxWorkers int, queueSize int) {
	JobQueue = make(chan Job, queueSize)
	d := NewDispatcher(maxWorkers)
	d.Run()
}
