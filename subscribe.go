package pubsub

type SubscribeInput struct {
	// Destination is the name of the destination to subscribe to.
	Destination string
	// AutoAck is a flag that indicates whether the subscription should automatically acknowledge
	// messages.
	AutoAck bool
}

// SetDestination sets the destination to subscribe to.
func (input *SubscribeInput) SetDestination(destination string) *SubscribeInput {
	input.Destination = destination
	return input
}

// SetAutoAck sets the flag that indicates whether the subscription should automatically
// acknowledge.
func (input *SubscribeInput) SetAutoAck(isAutoAck bool) *SubscribeInput {
	input.AutoAck = isAutoAck
	return input
}
