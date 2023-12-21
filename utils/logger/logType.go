package logger

type LogType string

const (
	ERR  LogType = "erreur"
	INFO LogType = "info"
)

type LogMessage struct {
	MsgType LogType `json:"type"`
	Donnee  any     `json:"data"`
}
