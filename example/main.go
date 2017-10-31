package main

import (
	"github.com/jackypanster/util"
	"time"
	"log"
)

func main() {
	start := time.Now()
	done := make(chan bool, 32)

	util.SetDebug(true)
	util.InitQueue(16, 1024)

	for i := 1; i <= 32; i ++ {
		item := i
		util.JobQueue <- util.Job{
			Do: func() error {
				log.Printf("sleep %d sec", item)
				time.Sleep(time.Second * time.Duration(item))
				done <- true
				return nil
			},
		}
		log.Printf("JobQueue %d", len(util.JobQueue))
	}

	for i := 0; i < 32; i ++ {
		<-done
	}
	log.Printf("complete %s", time.Now().Sub(start))
}
