package settings

import (
	pubSubSettings "github.com/cjburchell/pubsub/settings"
	"github.com/cjburchell/settings-go"
	log "github.com/cjburchell/uatu-go"
	"github.com/cjburchell/uatu-go/publishers"
)

// Get the log settings
func Get(settings settings.ISettings) log.Settings {
	return log.Settings{
		ServiceName:    settings.Get("LogServiceName", ""),
		MinLogLevel:    log.GetLogLevel(settings.Get("MinLogLevel", log.INFO.Text)),
		LogToConsole:   settings.GetBool("LogToConsole", true),
		HTTPSettings:   createHTTPSettings(settings.GetSection("http")),
		PubSubSettings: pubSubSettings.Get(settings.GetSection("pubSub")),
		UseHTTP:        settings.GetBool("LogUseHttp", false),
		UsePubSub:      settings.GetBool("LogUsePubSub", false),
	}
}

func createHTTPSettings(settings settings.ISettings) publishers.HTTPSettings {
	return publishers.HTTPSettings{
		Address: settings.Get("LogEndpoint", "http://logger:8082/log"),
		Token:   settings.Get("LogToken", "token"),
	}
}
