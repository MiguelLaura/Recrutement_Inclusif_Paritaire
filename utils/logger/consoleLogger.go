package logger

import (
	"fmt"
	"log"
)

type ConsoleLogger struct{}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (l *ConsoleLogger) Log(msg ...any) error {
	log.Print(msg...)
	return nil
}

func (l *ConsoleLogger) Logf(format string, v ...any) error {
	log.Printf(format, v...)
	return nil
}

func (l *ConsoleLogger) Err(msg ...any) error {
	log.Print(msg...)
	return nil
}

func (l *ConsoleLogger) LogType(logType LogType, msg ...any) error {
	return l.Log(append([]any{logType}, msg...)...)
}

func (l *ConsoleLogger) LogfType(logType LogType, format string, v ...any) error {
	return l.LogType(logType, fmt.Sprintf(format, v...))
}
