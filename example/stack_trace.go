package main

import (
	"errors"

	"github.com/jackypanster/util"
)

func test_trace() {
	err := errors.New("I don't know what to do yet")
	if err != nil {
		util.Panicf(util.Map{"key": "value", "cause": err}, "error occurs")
	}
}
