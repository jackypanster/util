package main

import (
	"sync"
	"fmt"
	"os"
	"strconv"
	"github.com/jackypanster/util"
)
var file = "log/log_%d.json"
var log *util.Log

func main() {
	pid := os.Getpid()
	fmt.Println("Own process identifier: ", strconv.Itoa(pid))

	log = util.NewProductLog(fmt.Sprintf(file,pid))
	defer log.Sync()
		 
	total := 100
	var wg sync.WaitGroup
	wg.Add(total)
	
	for i := 0; i < total; i++ {
		go test_log(i, &wg)
	}
	wg.Wait()
	fmt.Println("Main goroutine exit")
}
