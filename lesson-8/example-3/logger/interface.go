package logger

type IStringer interface {
	String() string
}

type ILogger interface {
	Info(message IStringer)
	Warn(message IStringer)
	Fatal(message IStringer)
	ReadLogs() []string
}
