package main

import (
	//"os"

	"log"
	"time"
	//"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	//f, err := os.OpenFile("log/log.json", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = append(cfg.OutputPaths, "./log/log.json")
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	//fmt.Printf("%#v\n", cfg)
	const url = "http://example.com"
	p := Person{
		Name: "jp",
		Age:  100,
	}
	// In most circumstances, use the SugaredLogger. It's 4-10x faster than most
	// other structured logging packages and has a familiar, loosely-typed API.
	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL.",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"person", p,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	//log.Panicf("err %s", errors.New("try"))
	start := time.Now()
	log.Printf("%s", time.Now().Sub(start))
}
