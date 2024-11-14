package pubsub

type Subscription interface {
	// Receive returns a channel that will receive messages from the subscription.
	Receive() <-chan Message
	// Unsubscribe unsubscribes from the subscription.
	Unsubscribe() error
}

type Message interface {
	// Ack acknowledges the message.
	Ack() error
	// Nack negatively acknowledges the message.
	Nack() error
	// Body returns the message body.
	Body() []byte
	// ContentType returns the content type of the message.
	ContentType() string
	// MessageID return the ID of the message.
	MessageID() string
}
