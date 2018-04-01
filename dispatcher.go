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
		// a job request has been received
		case job := <-JobQueue:
			// try to obtain a worker job channel that is available
			// this will block until a worker is idle
			go func(job Job) {
				jobChannel := <-d.workerPool
				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

func (d *Dispatcher) status() {
	for {
		if len(d.workerPool) == 0 {
			log.Printf("all workers are busy and items remain %d", len(JobQueue))
		}
		time.Sleep(time.Second * 2)
	}
}

func (d *Dispatcher) Workload() int {
	return len(d.workerPool)
}
