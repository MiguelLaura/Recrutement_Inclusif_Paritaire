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
	fichier *os.File
	ouvert  bool
}

func NewFileLogger(chemin string) (logger *FileLogger, err error) {
	logger = &FileLogger{}

	nomChemin, err := creerChemin(chemin)
	if err != nil {
		return nil, err
	}

	nomFichier := fmt.Sprintf("%s%s%s", nomChemin, loggerFileName, time.Now().Format(time.RFC3339))

	fichier, err := os.Create(nomFichier)
	if err != nil {
		return nil, err
	}

	// A la destruction du logger, Go va exécuter la function Fermer
	runtime.SetFinalizer(logger, Fermer)

	logger.fichier = fichier
	logger.ouvert = true

	return
}

func creerChemin(logFolder string) (chemin string, err error) {
	chemin = strings.TrimSpace(logFolder)

	if len(chemin) == 0 {
		return "./", nil
	}

	if len(chemin) > 0 && chemin[len(chemin)-1] != '/' {
		chemin += "/"
	}

	err = os.MkdirAll(chemin, os.ModePerm)

	return
}

func (l *FileLogger) Log(msg ...any) (err error) {
	now := time.Now()
	timeStr := fmt.Sprintf("%s %s ", now.Format(time.DateOnly), now.Format(time.TimeOnly))
	_, err = l.fichier.WriteString(timeStr + fmt.Sprint(msg...) + "\n")
	return
}

func (l *FileLogger) Logf(format string, v ...any) (err error) {
	err = l.Log(fmt.Sprintf(format, v...))
	return
}

func (l *FileLogger) Err(msg ...any) error {
	return l.LogType(ERR, msg...)
}

func (l *FileLogger) LogType(logType LogType, msg ...any) (err error) {

	err = l.Log(append([]any{formatLogType(logType)}, msg...)...)
	return
}

func (l *FileLogger) LogfType(logType LogType, format string, v ...any) (err error) {
	err = l.LogType(logType, fmt.Sprintf(format, v...))
	return
}

func (l *FileLogger) Fermer() {
	if l.ouvert {
		l.fichier.Close()
		l.ouvert = false
	}
}

func Fermer(l *FileLogger) {
	l.Fermer()
}
