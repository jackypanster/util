package main

import (
	"fmt"
	"os"

	"github.com/jackypanster/util"
)

func main() {
	f, err := os.OpenFile("log/log.json", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	util.CheckErrf(err, "fail to open file")
	util.SetOutput(f)
	util.SetDebug(true)
	defer f.Close()

	util.Debugf(util.Map{"key": "value"}, "%s %d", "testing", 1)
	util.Infof(util.Map{"key": "value"}, "%s %s", "testing", "more")
	util.Warnf(util.Map{"key": "value"}, "%s", "testing")
	util.Errorf(util.Map{"key": "value"}, "%s", "testing")

	util.Errorf(util.Map{"key": "value"}, "")

	type P struct {
		Name string
	}

	pStr := util.ToJsonString(P{Name: "jp"})

	fmt.Println(pStr)

	var p P
	util.ToInstance(pStr, &p)
	fmt.Printf("%+v", p)

	test_enq()
	test_deq()
}
