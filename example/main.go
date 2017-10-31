package main

import (
	"github.com/jackypanster/util"
	"log"
)

func main() {
	util.InitQueue(16, 1024)
	util.JobQueue <- util.Job{
		Do: func() error {
			log.Println("testing")
			return nil
		},
	}
}
