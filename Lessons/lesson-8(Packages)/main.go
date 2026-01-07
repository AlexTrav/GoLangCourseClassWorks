package main

import (
	"testPackage/counter"
	"testPackage/logger"
	_ "testPackage/pkg/printer"
)

func main() {
	log := logger.NewLogger()
	log.Info("Starting App")
	counter.Count()
	counter.Count()
	counter.Count()
	log.Warn("App is running with some warning")
	counter.Print()
	log.Info("App finished")
}
