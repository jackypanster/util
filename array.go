package util

import "strings"

func Uniq(src []string) []string {
	m := make(map[string]int)
	var dst []string

	for _, i := range src {
		item := strings.TrimSpace(i)
		if len(item) > 0 {
			m[item] = 0
		}
	}

	for key := range m {
		dst = append(dst, key)
	}
	return dst
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
