package log

import (
	"github.com/cjburchell/uatu-go/publishers"
)

// Settings for sending logs
type Settings struct {
	ServiceName  string
	MinLogLevel  Level
	LogToConsole bool
	UseNats      bool
	UseHttp      bool
	HTTPSettings publishers.HTTPSettings
	NatsSettings publishers.NatsSettings
}
