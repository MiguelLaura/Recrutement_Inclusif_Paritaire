package logger

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

const DEFAULT_BUFF_SIZE = 10

type SocketLogger struct {
	buffer []any
	ws     *websocket.Conn
}

func NewSocketLogger(ws *websocket.Conn, buffSize int) *SocketLogger {
	if buffSize <= 0 {
		buffSize = DEFAULT_BUFF_SIZE
	}
	return &SocketLogger{make([]any, 0, buffSize), ws}
}

func (sl *SocketLogger) Log(msg ...any) (err error) {
	return sl.LogType(INFO, msg...)
}

func (sl *SocketLogger) Logf(format string, v ...any) error {
	msg := fmt.Sprintf(format, v...)
	return sl.Log(msg)
}

func (sl *SocketLogger) Err(msg ...any) error {
	return sl.LogType(ERR, msg...)
}

func (sl *SocketLogger) LogType(logType LogType, msg ...any) error {
	msgReq := LogMessage{logType, msg}
	donnee, err := json.Marshal(msgReq)
	if err != nil {
		return err
	}
	err = sl.ws.WriteMessage(websocket.TextMessage, donnee)
	return err
}

func (sl *SocketLogger) LogfType(logType LogType, format string, v ...any) error {
	msg := fmt.Sprintf(format, v...)
	return sl.LogType(logType, msg)
}
