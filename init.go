package util

var JobQueue chan Job
var Debug = false

func init() {
	SetDebug(false) // default settings
}

func SetDebug(debug bool) {
	Debug = debug
}

func InitQueue(maxWorkers int, queueSize int) {
	JobQueue = make(chan Job, queueSize)
	d := NewDispatcher(maxWorkers)
	d.Run()
}
