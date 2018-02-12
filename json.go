package util

import (
	"encoding/json"
)

func ToJsonString(v interface{}) string {
	bs, err := json.Marshal(v)
	if err != nil {
		CheckErrf(err, "fail to marshal")
	}
	return string(bs)
}
