package util

import "log"

var JobQueue chan Job

func init() {
	log.Println("initialize package util")
}

func Init(maxWorkers int, length int) {
	JobQueue = make(chan Job, length)
	d := NewDispatcher(maxWorkers)
	d.Run()
}