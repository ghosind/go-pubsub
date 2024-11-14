package stomp

import (
	stomp3 "github.com/go-stomp/stomp/v3"
)

// StompMessage is a message for STOMP protocol.
type StompMessage struct {
	msg *stomp3.Message
}

// newStompMessage creates a new message for STOMP protocol.
func newStompMessage(msg *stomp3.Message) *StompMessage {
	stompMsg := new(StompMessage)
	stompMsg.msg = msg
	return stompMsg
}

// Ack acknowledges the message.
func (msg *StompMessage) Ack() error {
	return msg.msg.Conn.Ack(msg.msg)
}

// Nack negatively acknowledges the message.
func (msg *StompMessage) Nack() error {
	return msg.msg.Conn.Nack(msg.msg)
}

// Body returns the message body.
func (msg *StompMessage) Body() []byte {
	return msg.msg.Body
}

// ContentType returns the content type of the message.
func (msg *StompMessage) ContentType() string {
	return msg.msg.ContentType
}

// MessageID return the ID of the message.
func (msg *StompMessage) MessageID() string {
	return msg.msg.Header.Get("message-id")
}
