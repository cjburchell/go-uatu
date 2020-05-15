package log

import (
	"github.com/cjburchell/pubsub"
	"github.com/cjburchell/uatu-go/publishers"
)

// Settings for sending logs
type Settings struct {
	ServiceName    string
	MinLogLevel    Level
	LogToConsole   bool
	UsePubSub      bool
	UseHttp        bool
	HTTPSettings   publishers.HTTPSettings
	PubSubSettings pubsub.Settings
}
