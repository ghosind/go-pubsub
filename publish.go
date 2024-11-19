package pubsub

type PublishInput struct {
	// Destination is the name of the destination to publish to.
	Destination string
	// Body is the message body.
	Body []byte
	// ContentType is the content type of the message.
	ContentType string
	// MessageID is the id of the message.
	MessageID string
	// Persistent is a flag that indicates whether the message should be persisted.
	Persistent bool
	// Priority is the priority of the message.
	Priority int
}

// SetDestination sets the destination to publish to.
func (input *PublishInput) SetDestination(destination string) *PublishInput {
	input.Destination = destination
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

func (input *PublishInput) SetPriority(priority int) *PublishInput {
	input.Priority = priority
	return input
}
