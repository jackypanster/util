package main

import (
	"github.com/jackypanster/util"
	"log"
)

func main() {
	done := make(chan bool, 256)

	util.SetDebug(true)
	util.InitQueue(16, 1024)

	for i := 0; i < 256; i ++ {
		util.JobQueue <- util.Job{
			Do: func() error {
				log.Println("testing")
				done <- true
				return nil
			},
		}
	}

	for i := 0; i < 256; i ++ {
		<-done
	}
}
