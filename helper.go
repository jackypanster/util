package util

import (
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Panicf("[ERROR] %+v", err)
	}
}
