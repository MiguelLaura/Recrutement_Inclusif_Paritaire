package logger

import "sync"

type Loggers struct {
	sync.Mutex
	loggers []Logger
}

func NewLoggers() *Loggers {
	return &Loggers{sync.Mutex{}, make([]Logger, 0)}
}

func (l *Loggers) Log(msg ...any) {
	l.Lock()
	defer l.Unlock()
	for _, logger := range l.loggers {
		logger.Log(msg...)
	}
}

func (l *Loggers) Logf(format string, v ...any) {
	l.Lock()
	defer l.Unlock()
	for _, logger := range l.loggers {
		logger.Logf(format, v...)
	}
}

func (l *Loggers) Err(msg ...any) {
	l.Lock()
	defer l.Unlock()
	for _, logger := range l.loggers {
		logger.Err(msg...)
	}
}

func (l *Loggers) LogType(logType LogType, msg ...any) {
	l.Lock()
	defer l.Unlock()
	for _, logger := range l.loggers {
		logger.LogType(logType, msg...)
	}
}

func (l *Loggers) LogfType(logType LogType, format string, v ...any) {
	l.Lock()
	defer l.Unlock()
	for _, logger := range l.loggers {
		logger.LogfType(logType, format, v...)
	}
}

func (l *Loggers) AjouterLogger(nouveauLogger Logger) {
	l.Lock()
	defer l.Unlock()
	l.loggers = append(l.loggers, nouveauLogger)
}
