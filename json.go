package util

import (
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ToJsonString(v interface{}) string {
	bs, err := json.Marshal(v)
	CheckErrf(err, "fail to marshal")

	return string(bs)
}

func ToInstance(data string, v interface{}) {
	CheckStr(data, "data")
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		Errorf(Map{"error": err}, "")
	}
	CheckErrf(err, "fail to unmarshal")
}
