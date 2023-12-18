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

func NewFileLogger(path string) (logger *FileLogger, err error) {
	logger = &FileLogger{}

	pathName, err := createLogPath(path)
	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%s%s%s", pathName, loggerFileName, time.Now().Format(time.RFC3339))

	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	runtime.SetFinalizer(logger, Close)

	logger.file = file
	logger.opened = true

	return
}

func createLogPath(logFolder string) (pathName string, err error) {
	pathName = strings.TrimSpace(logFolder)

	if len(pathName) == 0 {
		return "./", nil
	}

	if len(pathName) > 0 && pathName[len(pathName)-1] != '/' {
		pathName += "/"
	}

	err = os.MkdirAll(pathName, os.ModePerm)

	return
}

func (l *FileLogger) Log(msg ...any) (err error) {
	_, err = l.file.WriteString(fmt.Sprint(msg...))
	return
}

func (l *FileLogger) Logf(format string, v ...any) (err error) {
	_, err = l.file.WriteString(fmt.Sprintf(format, v...))
	return
}

func (l *FileLogger) Err(msg ...any) (err error) {
	_, err = l.file.WriteString("ERR:" + fmt.Sprint(msg...))
	return
}

func (l *FileLogger) LogType(logType LogType, msg ...any) (err error) {
	err = l.Log(append([]any{logType}, msg...)...)
	return
}

func (l *FileLogger) LogfType(logType LogType, format string, v ...any) (err error) {
	err = l.LogType(logType, fmt.Sprintf(format, v...))
	return
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
