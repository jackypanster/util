package util

import (
	"log"
	"strconv"
	"strings"
)

func GetNumber(version string) int {
	if len(version) > 0 {
		tokens := strings.Split(version, ".")
		if len(tokens) >= 2 {
			first, err := strconv.Atoi(tokens[0])
			if err != nil {
				// handle error
				log.Printf("fail to parse version %s, %s", version, err.Error())
				return 0
			}
			second, err := strconv.Atoi(tokens[1])
			if err != nil {
				// handle error
				log.Printf("fail to parse version %s, %s", version, err.Error())
				return 0
			}
			return first*10 + second
		}
	}
	return 0
}
