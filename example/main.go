package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jackypanster/util"
	"time"
	"os"
	"fmt"
)

func main() {
	var filename = "log/logfile.log"
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	Formatter := new(log.JSONFormatter)
	//Formatter.TimestampFormat = "02-01-2006 15:04:05"
	log.SetFormatter(Formatter)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}
	defer f.Close()
	log.SetLevel(log.DebugLevel)

	start := time.Now()
	done := make(chan bool, 64)

	util.InitQueue(128, 64)

	for i := 1; i <= 64; i ++ {
		item := i
		util.JobQueue <- util.Job{
			Do: func() error {
				time.Sleep(time.Millisecond * time.Duration(item))
				done <- true
				return nil
			},
		}
	}

	for i := 0; i < 64; i ++ {
		<-done
	}
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Warnf("complete %s", time.Now().Sub(start))
}
