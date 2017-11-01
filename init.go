package util

var JobQueue chan Job

func InitQueue(maxWorkers int, queueSize int) {
	JobQueue = make(chan Job, queueSize)
	d := NewDispatcher(maxWorkers)
	d.Run()
}
