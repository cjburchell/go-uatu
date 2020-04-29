package publishers

type Publisher interface {
	// Publish message
	Publish(messageBites []byte) error
}
