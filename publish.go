package pubsub

type PublishInput struct {
	// Queue is the name of the queue to publish to.
	Queue string
	// Body is the message body.
	Body []byte
	// ContentType is the content type of the message.
	ContentType string
	// MessageID is the id of the message.
	MessageID string
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

// SetContentType sets the content type of the message.
func (input *PublishInput) SetContentType(contentType string) *PublishInput {
	input.ContentType = contentType
	return input
}

// SetMessageID sets the id of the message.
func (input *PublishInput) SetMessageID(messageId string) *PublishInput {
	input.MessageID = messageId
	return input
}

// SetPersistent sets the flag that indicates whether the message should be persisted.
func (input *PublishInput) SetPersistent(isPersistent bool) *PublishInput {
	input.Persistent = isPersistent
	return input
}
