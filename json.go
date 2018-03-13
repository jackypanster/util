package util

import (
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ToJsonString(v interface{}) (string, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		Errorf(Map{"error": err}, "unable to encode %#v", v)
		return "", err
	}
	return string(bs), nil
}

func ToInstance(data string, v interface{}) error {
	CheckStr(data, "data")
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		Errorf(Map{"error": err}, "unable to decode %s", data)
		return err
	}
	return nil
}
