package util

import (
  "strconv"
  "time"
)

func ConvertTimestamp(timestamp string) time.Time {
  i, err := strconv.ParseInt(timestamp, 10, 64)
  if err != nil {
    panic(err)
  }
  return time.Unix(0, i*int64(time.Millisecond))
}
