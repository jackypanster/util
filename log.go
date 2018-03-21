package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	*zap.SugaredLogger
}

func NewProductLog(file string) *Log {
	CheckStr(file, "file")
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths = append(cfg.OutputPaths, file)
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return &Log{
		logger.Sugar(),
	}
}

func NewDevLog() *Log {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return &Log{
		logger.Sugar(),
	}
}
