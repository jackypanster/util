package main

import (
	"github.com/jackypanster/util"
	"time"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("log/testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("This is a test log entry")

	start := time.Now()
	done := make(chan bool, 64)

	util.SetDebug(true)
	util.InitQueue(128, 64)

	for i := 1; i <= 64; i ++ {
		item := i
		util.JobQueue <- util.Job{
			Do: func() error {
				log.Printf("sleep %d sec", item)
				time.Sleep(time.Millisecond * time.Duration(item))
				done <- true
				return nil
			},
		}
		log.Printf("JobQueue %d", len(util.JobQueue))
	}

	for i := 0; i < 64; i ++ {
		<-done
	}
	log.Printf("complete %s", time.Now().Sub(start))
}
