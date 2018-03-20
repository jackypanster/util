package util

import (
	"io"

	"github.com/Gurpartap/logrus-stack"
	log "github.com/Sirupsen/logrus"
)

type Map map[string]interface{}

func init() {
	SetDebug(true)
	log.AddHook(logrus_stack.StandardHook())
}

func SetOutput(out io.Writer) {
	log.SetOutput(out)
}

func SetDebug(enable bool) {
	if enable {
		log.SetLevel(log.DebugLevel)
	}

	if enable {
		log.SetFormatter(&log.TextFormatter{
			ForceColors: true,
			TimestampFormat: "02-01-2006 15:04:05",
		})
	} else {
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "02-01-2006 15:04:05",
		})
	}
}

func Warnf(fields map[string]interface{}, format string, args ...interface{}) {
	log.WithFields(fields).Warnf(format, args...)
}

func Errorf(fields map[string]interface{}, format string, args ...interface{}) {
	log.WithFields(fields).Errorf(format, args...)
}

func Infof(fields map[string]interface{}, format string, args ...interface{}) {
	log.WithFields(fields).Infof(format, args...)
}

func Debugf(fields map[string]interface{}, format string, args ...interface{}) {
	log.WithFields(fields).Debugf(format, args...)
}

func Panicf(fields map[string]interface{}, format string, args ...interface{}) {
	log.WithFields(fields).Panicf(format, args...)
}
