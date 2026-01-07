package logger

import "fmt"

type FmtLogger struct{}

func (f *FmtLogger) Info(msg string) {
	fmt.Println("info: ", msg)
}

func (f *FmtLogger) Warn(msg string) {
	fmt.Println("warn: ", msg)
}
