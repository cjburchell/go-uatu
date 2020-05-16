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
		ServiceName:    settings.Get("ServiceName", ""),
		MinLogLevel:    log.GetLogLevel(settings.Get("MinLogLevel", log.INFO.Text)),
		LogToConsole:   settings.GetBool("LogToConsole", true),
		HTTPSettings:   createHTTPSettings(settings.GetSection("Http")),
		PubSubSettings: pubSubSettings.Get(settings.GetSection("PubSub")),
		UseHTTP:        settings.GetSection("Http").GetBool("Enabled", false),
		UsePubSub:      settings.GetSection("PubSub").GetBool("Enabled", false),
	}
}

func createHTTPSettings(settings settings.ISettings) publishers.HTTPSettings {
	return publishers.HTTPSettings{
		Address: settings.Get("Endpoint", "http://logger:8082/log"),
		Token:   settings.Get("Token", "token"),
	}
}
