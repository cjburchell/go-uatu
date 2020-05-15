package publishers

import (
	"bytes"
	"fmt"
	"net/http"
)

// HTTPSettings struct
type HTTPSettings struct {
	Address string
	Token   string
}

type httpPublisher struct {
	restClient *http.Client
	settings   HTTPSettings
}

func (publisher httpPublisher) Publish(messageBites []byte) error {
	req, err := http.NewRequest("POST", publisher.settings.Address, bytes.NewBuffer(messageBites))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("APIKEY %s", publisher.settings.Token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := publisher.restClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unable to send log to %s(%d)", publisher.settings.Address, resp.StatusCode)
	}

	return nil
}

// SetupHTTP sets up the http client
func SetupHTTP(newSettings HTTPSettings) Publisher {
	restClient := &http.Client{}
	return httpPublisher{restClient: restClient, settings: newSettings}
}
