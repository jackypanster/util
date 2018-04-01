package main

import (
	"fmt"

	"github.com/jackypanster/util"
)

var file = "log/log_%d.json"
var log *util.Log

func main() {
	var m map[string]string
	fmt.Println(len(m))
	fmt.Println(m == nil)

	var s string
	fmt.Println(len(s))

	var v interface{}
	fmt.Println(v == nil)
}
