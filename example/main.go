package main

import (
	"log"
	"time"
)

func main() {
	log.Println("main")
	start := time.Now()
	time.Sleep(time.Millisecond * 100)
	cost := time.Now().Sub(start)
	if cost > time.Second {
		log.Println()
	}
}
