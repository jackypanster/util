package main

import (
	"encoding/json"
	"log"
)

func main() {
	alert := make(map[string]string)
	alert["title"] = "v1"
	alert["body"] = "v2"
	b, _ := json.Marshal(alert)
	log.Println(string(b))
}
