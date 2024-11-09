package stomp

import (
	stomp3 "github.com/go-stomp/stomp/v3"
)

type StompMessage struct {
	msg *stomp3.Message
}

func newStompMessage(msg *stomp3.Message) *StompMessage {
	stompMsg := new(StompMessage)
	stompMsg.msg = msg
	return stompMsg
}

func (msg *StompMessage) Ack() error {
	return msg.msg.Conn.Ack(msg.msg)
}

func (msg *StompMessage) Nack() error {
	return msg.msg.Conn.Nack(msg.msg)
}

func (msg *StompMessage) Body() []byte {
	return msg.msg.Body
}
