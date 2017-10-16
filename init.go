package util

import "log"

var JobQueue chan Job

func init() {
	log.Println("initialize package util")
}

func InitQueue(maxWorkers int, queueSize int) {
	JobQueue = make(chan Job, queueSize)
	d := NewDispatcher(maxWorkers)
	d.Run()
}
