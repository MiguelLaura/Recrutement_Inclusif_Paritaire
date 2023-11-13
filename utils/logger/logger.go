package logger

type Logger interface {
	Log(msg ...any)
	Err(msg ...any)
}
