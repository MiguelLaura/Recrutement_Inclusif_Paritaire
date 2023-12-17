package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

const loggerFileName = "log"

type FileLogger struct {
	file   *os.File
	opened bool
}

func NewFileLogger(logFile string) (logger *FileLogger, err error) {
	logger = &FileLogger{}

	fileName := strings.TrimSpace(logFile)

	if fileName == "" {
		fileName = loggerFileName + time.Now().Format(time.RFC3339)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	runtime.SetFinalizer(logger, Close)

	logger.file = file
	logger.opened = true

	return
}

func (l *FileLogger) Log(msg ...any) {
	l.file.WriteString(fmt.Sprint(msg...))
}

func (l *FileLogger) Err(msg ...any) {
	l.file.WriteString("ERR:" + fmt.Sprint(msg...))
}

func (l *FileLogger) Close() {
	if l.opened {
		l.file.Close()
		l.opened = false
	}
}

func Close(l *FileLogger) {
	l.Close()
}
