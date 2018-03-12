package util

import (
	log "github.com/Sirupsen/logrus"
)

func Warnf(fields map[string]interface{}, format string, args ...interface{}) {
	if len(fields) != 0 {
		log.WithFields(fields).Warnf(format, args)
	} else {
		log.Warnf(format, args)
	}
}
