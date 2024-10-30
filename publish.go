package pubsub

type PublishInput struct {
	// Queue is the name of the queue to publish to.
	Queue string
	// Body is the message body.
	Body []byte
	// Persistent is a flag that indicates whether the message should be persisted.
	Persistent bool
}

// SetQueue sets the queue to publish to.
func (input *PublishInput) SetQueue(queue string) *PublishInput {
	input.Queue = queue
	return input
}

// SetBody sets the message body.
func (input *PublishInput) SetBody(body []byte) *PublishInput {
	input.Body = body
	return input
}

// SetPersistent sets the flag that indicates whether the message should be persisted.
func (input *PublishInput) SetPersistent(isPersistent bool) *PublishInput {
	input.Persistent = isPersistent
	return input
}
