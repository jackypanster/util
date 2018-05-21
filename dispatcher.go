package util

import (
	"log"
	"time"
)

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	workerPool chan chan Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{workerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.workerPool)
		worker.SetId(i)
		worker.Start()
	}
	go d.dispatch()
	go d.status()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			//go func(job Job) {
			// try to obtain a worker job channel that is available
			// this will block until a worker is idle
			jobChannel := <-d.workerPool

			// dispatch the job to the worker job channel
			jobChannel <- job
			//}(job)
		}
	}
}

func (d *Dispatcher) status() {
	for {
		log.Printf("%d workers available, %d jobs in the buffer", len(d.workerPool), len(JobQueue))
		time.Sleep(time.Minute)
	}
}

func (d *Dispatcher) Workload() int {
	return len(d.workerPool)
}
