package main

import (
	"github.com/jackypanster/util"
	"log"
)

func main() {
	done := make(chan bool)

	util.SetDebug(false)
	util.InitQueue(16, 1024)
	util.JobQueue <- util.Job{
		Do: func() error {
			log.Println("testing")
			done <- true
			return nil
		},
	}

	<-done
}
