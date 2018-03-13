package main

import (
	"fmt"

	"github.com/jackypanster/util"
	"github.com/mitchellh/mapstructure"
)

var pool = util.GetRedisPool("127.0.0.1:6379")

var redisService = util.NewRedisService(pool, "testQ")

var taskQ = util.NewTaskService(redisService)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func test_enq() {
	p := Person{
		Name: "jp",
		Age:  100,
	}
	err := taskQ.Enq("133", p, 234)
	util.CheckErr(err)
}

func test_deq() {
	task, err := taskQ.Deq()
	util.CheckErr(err)
	p, ok := task.Content.(map[string]interface{})
	fmt.Printf("%#v\n", task)
	if ok {
		var result Person
		mapstructure.Decode(p, &result)
		fmt.Printf("%s, %d\n", result.Name, result.Age)
	} else {
		fmt.Println("cast not ok")
	}
}
