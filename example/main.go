package main

import (
	"github.com/jackypanster/util"
	"time"
	"log"
)

func main() {
	start := time.Now()
	done := make(chan bool, 128)

	util.SetDebug(true)
	util.InitQueue(8, 64)

	for i := 1; i <= 128; i ++ {
		item := i
		util.JobQueue <- util.Job{
			Do: func() error {
				log.Printf("sleep %d sec", item)
				time.Sleep(time.Microsecond * time.Duration(item))
				done <- true
				return nil
			},
		}
		log.Printf("JobQueue %d", len(util.JobQueue))
	}

	for i := 0; i < 128; i ++ {
		<-done
	}
	log.Printf("complete %s", time.Now().Sub(start))
}
