package util

import (
	"io"

	log "github.com/Sirupsen/logrus"
)

type Map map[string]interface{}

func init() {
	formatter := new(log.JSONFormatter)
	formatter.TimestampFormat = "02-01-2006 15:04:05"
	log.SetFormatter(formatter)
}

func SetOutput(out io.Writer) {
	log.SetOutput(out)
}

func SetDebug(enable bool) {
	if enable {
		log.SetLevel(log.DebugLevel)
	}
}

func Warnf(fields map[string]interface{}, format string, args ...interface{}) {
	if len(fields) != 0 {
		log.WithFields(fields).Warnf(format, args...)
	} else {
		log.Warnf(format, args...)
	}
}

func Errorf(fields map[string]interface{}, format string, args ...interface{}) {
	if len(fields) != 0 {
		log.WithFields(fields).Errorf(format, args...)
	} else {
		log.Errorf(format, args...)
	}
}

func Infof(fields map[string]interface{}, format string, args ...interface{}) {
	if len(fields) != 0 {
		log.WithFields(fields).Infof(format, args...)
	} else {
		log.Infof(format, args...)
	}
}

func Debugf(fields map[string]interface{}, format string, args ...interface{}) {
	if len(fields) != 0 {
		log.WithFields(fields).Debugf(format, args...)
	} else {
		log.Debugf(format, args...)
	}
}

func Panicf(fields map[string]interface{}, format string, args ...interface{}) {
	if len(fields) != 0 {
		log.WithFields(fields).Panicf(format, args...)
	} else {
		log.Panicf(format, args...)
	}
}
