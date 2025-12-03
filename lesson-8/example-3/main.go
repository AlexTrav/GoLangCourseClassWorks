package main

import (
	"exapmle-3/logger"
	"fmt"
	"time"
)

type LogMessage struct {
	UserId    int
	PaymentId int
	Message   string
	Formated  string
}

func (l *LogMessage) String() string {
	l.Formated = fmt.Sprintf(
		"%s: PaymentId: %d, UserId: %d, Message: %s",
		time.Now().Format("2006-01-02 15:04:05"),
		l.PaymentId,
		l.UserId,
		l.Message,
	)
	return l.Formated
}

func main() {
	var l logger.ILogger = &logger.Logger{}

	userId := 123
	paymentId := 222

	l.Info(&LogMessage{
		UserId:    userId,
		PaymentId: paymentId,
		Message:   "Test info: Payment Created",
	})

	l.Warn(&LogMessage{
		UserId:    userId,
		PaymentId: paymentId,
		Message:   "Test warn: Payment not created",
	})

	l.Fatal(&LogMessage{
		UserId:    userId,
		PaymentId: paymentId,
		Message:   "Test fatal: Payment not created; Fatal error",
	})

	for i, s := range l.ReadLogs() {
		fmt.Printf("Message %d: %s\n", i+1, s)
	}
}

/* Вывод из консоли:
[INFO] [2025-12-03 20:18:11: PaymentId: 222, UserId: 123, Message: Test info: Payment Created]
[WARN] [2025-12-03 20:18:11: PaymentId: 222, UserId: 123, Message: Test warn: Payment not created]
[FATAL] [2025-12-03 20:18:11: PaymentId: 222, UserId: 123, Message: Test fatal: Payment not created; Fatal error]
Message 1: 2025-12-03 20:18:11: PaymentId: 222, UserId: 123, Message: Test info: Payment Created
Message 2: 2025-12-03 20:18:11: PaymentId: 222, UserId: 123, Message: Test warn: Payment not created
Message 3: 2025-12-03 20:18:11: PaymentId: 222, UserId: 123, Message: Test fatal: Payment not created; Error
*/
