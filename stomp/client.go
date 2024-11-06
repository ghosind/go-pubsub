package stomp

import (
	stomp3 "github.com/go-stomp/stomp/v3"
)

// StompClient is a Pub-Sub client for STOMP protocol.
type StompClient struct {
	// conn is the STOMP connection.
	conn *stomp3.Conn
}
