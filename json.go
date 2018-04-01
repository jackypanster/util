package util

import (
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ToJsonString(v interface{}) (string, error) {
	CheckNil(v, "v")
	bs, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func ToInstance(str string, v interface{}) error {
	CheckStr(str, "str")
	CheckNil(v, "v")
	err := json.Unmarshal([]byte(str), v)
	if err != nil {
		return err
	}
	return nil
}
