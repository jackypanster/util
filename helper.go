package util

import (
	"log"
	"strings"
)

func CheckArray(values []string, name string) []string {
	var results []string
	if len(values) == 0 {
		log.Panicf("%s should not be empty", name)
	}
	m := make(map[string]int)
	for _, value := range values {
		str := strings.TrimSpace(value)
		if len(str) > 0 {
			m[str] = 0
		}
	}
	if len(m) == 0 {
		log.Panicf("%s should not be empty", name)
	}
	for key := range m {
		results = append(results, key)
	}
	return results
}

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