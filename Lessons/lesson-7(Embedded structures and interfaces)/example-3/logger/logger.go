package logger

import "fmt"

type Logger struct {
	infoData  []string
	warnData  []string
	fatalData []string
}

func (l *Logger) Info(message IStringer) {
	l.infoData = append(l.infoData, message.String())
	fmt.Println("[INFO]", l.infoData)
}

func (l *Logger) Warn(message IStringer) {
	l.warnData = append(l.warnData, message.String())
	fmt.Println("[WARN]", l.warnData)
}

func (l *Logger) Fatal(message IStringer) {
	l.fatalData = append(l.fatalData, message.String())
	fmt.Println("[FATAL]", l.fatalData)
}

func (l *Logger) ReadLogs() []string {
	var logs []string
	logs = append(logs, l.infoData...)
	logs = append(logs, l.warnData...)
	logs = append(logs, l.fatalData...)
	return logs
}
