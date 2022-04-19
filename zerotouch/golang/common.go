package main

import "go.uber.org/zap"

var logger = getLogger()

func getLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger
}
