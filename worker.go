package util

import (
	"time"
)

// Job represents the job to be run
type Job struct {
	Data   interface{}
	Do     func() error
	Handle func(any interface{}) error
}

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
	id         int
}

func NewWorker(workerPool chan chan Job) *Worker {
	return &Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w *Worker) SetId(num int) {
	w.id = num
}

// Start method starts the run loop for the worker, listening for a quit channel in case we need to stop it
func (w *Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				start := time.Now()
				if err := job.Do(); err != nil {
					Errorf(Map{"worker": w.id, "error": err}, "")
				}
				if err := job.Handle(job.Data); err != nil {
					Errorf(Map{"worker": w.id, "error": err, "data": job.Data}, "")
				}
				cost := time.Now().Sub(start)
				if cost > time.Minute {
					Warnf(nil, "worker#%d spends %s", w.id, cost)
				}
			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop method signals the worker to stop listening for work requests
func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
