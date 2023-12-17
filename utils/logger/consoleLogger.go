package logger

import (
	"fmt"
	"log"
)

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

func (l *ConsoleLogger) LogType(logType LogType, msg ...any) {
	l.Log(append([]any{logType}, msg...)...)
}

func (l *ConsoleLogger) LogfType(logType LogType, format string, v ...any) {
	l.LogType(logType, fmt.Sprintf(format, v...))
}

func (l *ConsoleLogger) ErrType(logType LogType, msg ...any) {
	l.Err(append([]any{logType}, msg...)...)
}
