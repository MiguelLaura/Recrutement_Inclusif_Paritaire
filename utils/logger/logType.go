package logger

type LogType string

const (
	ERR  LogType = "Error"
	INFO LogType = "Info"
)

type LogMessage struct {
	MsgType LogType `json:"type"`
	Data    any     `json:"data"`
}
