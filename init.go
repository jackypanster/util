package util

import "log"

var JobQueue chan Job

func init() {
	log.Println("initialize package util")
}

func InitQueue(maxWorkers int, size int) {
	JobQueue = make(chan Job, size)
	d := NewDispatcher(maxWorkers)
	d.Run()
}