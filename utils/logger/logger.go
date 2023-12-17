package logger

type Logger interface {
	Log(msg ...any)
	Logf(format string, v ...any)
	Err(msg ...any)

	LogType(logType LogType, msg ...any)
	LogfType(logType LogType, format string, v ...any)
	ErrType(logType LogType, msg ...any)
}
