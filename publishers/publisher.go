package publishers

type Publisher interface {
	Publish(messageBites []byte) error
}
