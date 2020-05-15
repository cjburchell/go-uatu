package publishers

// Publisher interface
type Publisher interface {
	// Publish message
	Publish(messageBites []byte) error
}
