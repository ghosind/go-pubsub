package pubsub

type SubscribeInput struct {
	// Queue is the name of the queue to subscribe to.
	Queue string
	// AutoAck is a flag that indicates whether the subscription should automatically acknowledge
	// messages.
	AutoAck bool
}

// SetQueue sets the queue to subscribe to.
func (input *SubscribeInput) SetQueue(queue string) *SubscribeInput {
	input.Queue = queue
	return input
}

// SetAutoAck sets the flag that indicates whether the subscription should automatically
// acknowledge.
func (input *SubscribeInput) SetAutoAck(isAutoAck bool) *SubscribeInput {
	input.AutoAck = isAutoAck
	return input
}
