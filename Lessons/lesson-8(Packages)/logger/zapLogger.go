package logger

import (
	"time"

	"go.uber.org/zap"
)

type ZapLogger struct {
	zapLogger *zap.SugaredLogger
}

func (z *ZapLogger) Info(msg string) {
	z.zapLogger.Info(msg)
}

func (z *ZapLogger) Warn(msg string) {
	z.zapLogger.Warn(msg)
}

func NewZapLogger() ILogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	return &ZapLogger{zapLogger: sugar.With("time", time.Now())}
}
