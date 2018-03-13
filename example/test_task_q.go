package main

import (
	"github.com/jackypanster/util"
)

var pool = util.GetRedisPool("127.0.0.1:6379")

var redisService = util.NewRedisService(pool, "testQ")

var taskQ = util.NewTaskService(redisService)

func test_enq() {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{
		Name: "jp",
		Age:  100,
	}
	taskQ.Enq()
}
