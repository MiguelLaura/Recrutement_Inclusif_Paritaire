package logger

import "log"

type ConsoleLogger struct{}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (l *ConsoleLogger) Log(msg ...any) {
	log.Print(msg...)
}

func (l *ConsoleLogger) Logf(format string, v ...any) {
	log.Printf(format, v...)
}

func (l *ConsoleLogger) Err(msg ...any) {
	log.Print(msg...)
}
