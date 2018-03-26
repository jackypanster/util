package main

import (
	"fmt"
	"log"
)

var chineseDict = map[string]string{
	"云": "雲",
}

func test_str() {
	str := "Hello, 云"
	/*fmt.Println(utf8.RuneCountInString(str))

	for idx, c := range str {
		fmt.Printf("%d: %s\n", idx, string(c))
	}

	yes := contains("云")
	fmt.Println(yes)

	val := convert("云")
	fmt.Println(val)

	val = convert("abcd")
	fmt.Println(val)*/

	fmt.Println(convert(str))
}

func contains(s string) bool {
	_, ok := chineseDict[s]
	return ok
}

func convert(str string) string {
	text := ""
	for idx, c := range str {
		log.Printf("%d, %s", idx, string(c))
		token := translate(string(c))
		text += token
	}
	return text
}

func translate(s string) string {
	if v, ok := chineseDict[s]; ok {
		return v
	}
	return s
}
