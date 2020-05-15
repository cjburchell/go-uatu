package settings

import (
	"github.com/cjburchell/settings-go"
	log "github.com/cjburchell/uatu-go"
	"github.com/cjburchell/uatu-go/publishers"
)

func Get(settings settings.ISettings) log.Settings {
	return log.Settings{
		ServiceName:  settings.Get("LOG_SERVICE_NAME", ""),
		MinLogLevel:  log.GetLogLevel(settings.Get("LOG_LEVEL", log.INFO.Text)),
		LogToConsole: settings.GetBool("LOG_CONSOLE", true),
		HTTPSettings: createHTTPSettings(settings.GetSection("http")),
		NatsSettings: createNatsSettings(settings.GetSection("nats")),
		UseHttp:      settings.GetSection("http").GetBool("LOG_USE_REST", false),
		UseNats:      settings.GetSection("nats").GetBool("LOG_USE_NATS", true),
	}
}

func createHTTPSettings(settings settings.ISettings) publishers.HTTPSettings {
	return publishers.HTTPSettings{
		Address: settings.Get("LOG_REST_URL", "http://logger:8082/log"),
		Token:   settings.Get("LOG_REST_TOKEN", "token"),
	}
}

func createNatsSettings(settings settings.ISettings) publishers.NatsSettings {
	return publishers.NatsSettings{
		URL:      settings.Get("LOG_NATS_URL", "tcp://nats:4222"),
		Token:    settings.Get("LOG_NATS_TOKEN", "token"),
		User:     settings.Get("LOG_NATS_USER", "admin"),
		Password: settings.Get("LOG_NATS_PASSWORD", "password"),
	}
}
