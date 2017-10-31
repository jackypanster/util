package util

import "log"

// Job represents the job to be run
type Job struct {
	Do func() error
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
				if Debug {
					log.Printf("[Worker %d] starts", w.id)
				}
				if err := job.Do(); err != nil {
					log.Printf("[ERROR] %s\n", err.Error())
				}
				if Debug {
					log.Printf("[Worker %d] ends", w.id)
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
