package publishers

import (
	"context"

	"github.com/cjburchell/pubsub"
)

type pubSubPublisher struct {
	connection pubsub.IPubSub
}

// Publish message
func (publisher pubSubPublisher) Publish(messageBites []byte) error {
	return publisher.connection.Publish(context.Background(), "logs", messageBites)
}

// Setup Pub Sub
func SetupPubSub(newSettings pubsub.Settings) (Publisher, error) {
	connection, err := pubsub.Create(context.Background(), newSettings)
	if err != nil {
		return nil, err
	}

	return pubSubPublisher{connection: connection}, nil
}
