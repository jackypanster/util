package util

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func init() {
	formatter := new(log.JSONFormatter)
	formatter.TimestampFormat = "02-01-2006 15:04:05"
	log.SetFormatter(formatter)
}

func SetOutput(filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	CheckErrf(err, "fail to open file")
	log.SetOutput(f)
	defer f.Close()
}

func Warnf(fields map[string]interface{}, format string, args ...interface{}) {
	if len(fields) != 0 {
		log.WithFields(fields).Warnf(format, args)
	} else {
		log.Warnf(format, args)
	}
}
