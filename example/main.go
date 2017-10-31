package main

import (
	"github.com/jackypanster/util"
	"time"
	"log"
)

func main() {
	start := time.Now()
	done := make(chan bool, 65536)

	util.SetDebug(true)
	util.InitQueue(8, 128)

	for i := 1; i <= 65536; i ++ {
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

	for i := 0; i < 65536; i ++ {
		<-done
	}
	log.Printf("complete %s", time.Now().Sub(start))
}
