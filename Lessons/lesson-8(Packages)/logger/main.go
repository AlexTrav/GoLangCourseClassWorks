package logger

import (
	"fmt"
	"testPackage/config"
)

func init() {
	fmt.Println("logger initialized")
}

const (
	DriverFmt = "fmt"
	DriverZap = "zap"
)

func NewLogger() ILogger {
	switch config.Cfg.Logger.Driver {
	case DriverFmt:
		return &FmtLogger{}
	case DriverZap:
		return NewZapLogger()
	}
	return nil
}
