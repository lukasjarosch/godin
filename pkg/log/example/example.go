package main

import "github.com/lukasjarosch/godin/pkg/log"

func main() {

	logger := log.New()
	logger.Debug("test", "foo", "bar")
	logger.Notice("test", "foo", "bar")
	logger.Info("test", "foo", "bar")
	logger.Warning("test", "foo", "bar")
	logger.Error("test", "foo", "bar")
	logger.Alert("test", "foo", "bar")
	logger.Critical("test", "foo", "bar")
	logger.Emergency("test", "foo", "bar")
}
