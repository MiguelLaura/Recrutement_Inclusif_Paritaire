package logger

import "fmt"

type Logger interface {
	Log(msg ...any) error
	Logf(format string, v ...any) error
	Err(msg ...any) error

	LogType(logType LogType, msg ...any) error
	LogfType(logType LogType, format string, v ...any) error
}

func formatLogType(logType LogType) LogType {
	return LogType(fmt.Sprintf("[%s] ", logType))
}
