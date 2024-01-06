package logger

import (
	"fmt"
	"time"
)

type ConsoleLogger struct{}

func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

func (l *ConsoleLogger) Log(msg ...any) error {
	now := time.Now()
	timeStr := fmt.Sprintf("%s %s", now.Format(time.DateOnly), now.Format(time.TimeOnly))
	fmt.Println(append([]any{timeStr}, msg...)...)
	return nil
}

func (l *ConsoleLogger) Logf(format string, v ...any) error {
	l.Log(fmt.Sprintf(format, v...))
	return nil
}

func (l *ConsoleLogger) Err(msg ...any) error {
	return l.LogType(ERR, msg...)
}

func (l *ConsoleLogger) LogType(logType LogType, msg ...any) error {
	return l.Log(append([]any{formatLogType(logType)}, msg...)...)
}

func (l *ConsoleLogger) LogfType(logType LogType, format string, v ...any) error {
	return l.LogType(logType, fmt.Sprintf(format, v...))
}
