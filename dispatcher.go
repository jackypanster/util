package util

import (
	log "github.com/Sirupsen/logrus"
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
}

func (d *Dispatcher) dispatch() {
	for {
		log.Debugf("jobs left %d; available workers remain %d", len(JobQueue), len(d.workerPool))
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
