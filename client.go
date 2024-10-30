package pubsub

import "context"

type Client interface {
	// Connect connects to the message broker.
	Connect() error
	// Publish publishes a message to a queue.
	Publish(*PublishInput) error
	// PublishWithContext publishes a message to a queue with a context.
	PublishWithContext(context.Context, *PublishInput) error
	// Subscribe subscribes to a queue.
	Subscribe(*SubscribeInput) (Subscription, error)
	// SubscribeWithContext subscribes to a queue with a context.
	SubscribeWithContext(context.Context, *SubscribeInput) (Subscription, error)
	// Close closes the connection to the message broker.
	Close() error
}
