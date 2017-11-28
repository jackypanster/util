package util

import (
	"strconv"
	"time"
)

func ConvertTimestamp(timestamp string) time.Time {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	CheckErr(err)
	return time.Unix(0, i*int64(time.Millisecond))
}

func ConvertDateString(date string) int64 {
	t, err := time.Parse("2006-01-02", date)
	CheckErr(err)
	return t.Unix()
}
