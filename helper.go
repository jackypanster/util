package util

import (
	"log"
	"strings"
)

func CheckStr(value string, name string) string {
	str := strings.TrimSpace(value)
	if len(str) == 0 {
		log.Panicf("%s should not be empty", name)
	}
	return str
}

func CheckErr(err error) {
	if err != nil {
		log.Panicf("[ERROR] %+v", err)
	}
}
