package util

import (
	"fmt"
	"log"
	"strings"
)

func CheckMap(m map[string]string, name string) map[string]string {
	result := make(map[string]string)

	CheckCondition(len(m) == 0, fmt.Sprintf("%s should not be empty", name))

	for k, v := range m {
		key := CheckStr(k, "key")
		val := CheckStr(v, "val")
		result[key] = val
	}

	CheckCondition(len(result) == 0, fmt.Sprintf("%s should not be empty", name))

	return result
}

func CheckMapInf(m map[string]interface{}, name string) map[string]interface{} {
	result := make(map[string]interface{})

	CheckCondition(len(m) == 0, fmt.Sprintf("%s should not be empty", name))

	for k, v := range m {
		key := CheckStr(k, "key")
		CheckNil(v, "val")
		result[key] = v
	}

	CheckCondition(len(result) == 0, fmt.Sprintf("%s should not be empty", name))

	return result
}

func CheckArray(values []string, name string) []string {
	var results []string

	CheckCondition(len(values) == 0, fmt.Sprintf("%s should not be empty", name))

	m := make(map[string]int)
	for _, value := range values {
		str := CheckStr(value, "value")
		m[str] = 0
	}

	CheckCondition(len(m) == 0, fmt.Sprintf("%s should not be empty", name))

	for key := range m {
		results = append(results, key)
	}
	return results
}

func CheckNil(v interface{}, name string) {
	CheckCondition(v == nil, fmt.Sprintf("%s should not be nil", name))
}

func CheckStr(value string, name string) string {
	str := strings.TrimSpace(value)
	CheckCondition(len(str) == 0, fmt.Sprintf("%s should not be empty", name))
	return str
}

func CheckCondition(condition bool, description string) {
	if condition {
		log.Panicf("%s", description)
	}
}

func CheckErr(err error) {
	CheckErrf(err, "error occurs")
}

func CheckErrf(err error, description string) {
	if err != nil {
		log.Panicf("%s, %s", description, err.Error())
	}
}
