package main

import (
	"github.com/jackypanster/util"
	"time"
	"log"
)

func main() {
	done := make(chan bool, 16)

	util.SetDebug(true)
	util.InitQueue(8, 1024)

	/*for i := 1; i <= 16; i ++ {
		util.JobQueue <- util.Job{
			Do: func() error {
				log.Printf("sleep %d sec", i)
				time.Sleep(time.Second * time.Duration(i))
				done <- true
				return nil
			},
		}
	}*/

	for i := 0; i < 16; i ++ {
		<-done
	}
}
