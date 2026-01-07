package config

import (
	"fmt"
	"os"
	"strconv"
	"testPackage/default/values"
)

type Counter struct {
	Interval int
}

type Logger struct {
	Driver string
}

type Config struct {
	Counter *Counter
	Logger  *Logger
}

var Cfg *Config

func init() {
	interval, err := strconv.Atoi(os.Getenv("INTERVAL"))
	if err != nil {
		interval = defaultValues.DefaultInterval
		fmt.Println("set interval to default value")
	}

	logDriver := os.Getenv("LOGGER_DRIVER")
	if logDriver == "" {
		logDriver = defaultValues.DefaultLogDriver
	}

	Cfg = &Config{Counter: &Counter{Interval: interval}, Logger: &Logger{Driver: logDriver}}
	fmt.Println("config initialized")
}
