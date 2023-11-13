package logger

import "log"

type ConsoleLogger struct{}

func (l *ConsoleLogger) Log(msg ...any) {
	log.Print(msg...)
}

func (l *ConsoleLogger) Err(msg ...any) {
	log.Print(msg...)
}
