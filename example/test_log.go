package main

import (
	"sync"
)

func test_log(i int, wg *sync.WaitGroup) {	
	defer wg.Done()
	log.Infof("%d", i)
}