package util

import (
	log "github.com/Sirupsen/logrus"
)

func init() {
	formatter := new(log.JSONFormatter)
	formatter.TimestampFormat = "02-01-2006 15:04:05"
	log.SetFormatter(formatter)
}

func Warnf(fields map[string]interface{}, format string, args ...interface{}) {
	if len(fields) != 0 {
		log.WithFields(fields).Warnf(format, args)
	} else {
		log.Warnf(format, args)
	}
}
