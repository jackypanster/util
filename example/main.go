package main

import (
	"encoding/json"
	"log"

	"github.com/jackypanster/util"
)

func main() {
	alert := make(map[string]string)
	alert["title"] = "v1"
	alert["body"] = "v2"
	b, _ := json.Marshal(alert)
	log.Println(string(b))

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{"j", 2}
	str := util.ToJsonString(p)
	log.Println(str)

	var person Person
	util.ToStructure(str, &person)
	log.Printf("%+v", person)
}
