package logger

type Logger interface {
	Log(msg ...any)
	Logf(format string, v ...any)
	Err(msg ...any)
}
