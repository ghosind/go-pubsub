package stomp

import (
	"context"
	"sync"

	"github.com/ghosind/go-pubsub"
	stomp3 "github.com/go-stomp/stomp/v3"
	"github.com/go-stomp/stomp/v3/frame"
)

// StompClient is a Pub-Sub client for STOMP protocol.
type StompClient struct {
	// conn is the STOMP connection.
	conn *stomp3.Conn

	// address is the address of the message broker.
	address string
	// username is the username for authentication.
	username string
	// password is the password for authentication.
	password string

	// connMutex is the mutex for connection.
	connMutex *sync.RWMutex
}

// Connect connects to the message broker.
func (cli *StompClient) Connect() error {
	cli.connMutex.Lock()
	defer cli.connMutex.Unlock()

	if cli.conn != nil {
		return nil
	}

	conn, err := stomp3.Dial(
		"tcp",
		cli.address,
		stomp3.ConnOpt.Login(cli.username, cli.password),
	)
	if err != nil {
		return err
	}

	cli.conn = conn
	return nil
}

// Publish publishes a message to a queue.
func (cli *StompClient) Publish(input *pubsub.PublishInput) error {
	return cli.PublishWithContext(context.Background(), input)
}

// PublishWithContext publishes a message to a queue with a context.
func (cli *StompClient) PublishWithContext(ctx context.Context, input *pubsub.PublishInput) error {
	cli.connMutex.RLock()
	defer cli.connMutex.RUnlock()

	if cli.conn == nil {
		if err := cli.Connect(); err != nil {
			return err
		}
	}

	opts := cli.makeSendOptions(input)
	contentType := input.ContentType
	if contentType == "" {
		contentType = "text/plain"
	}

	err := cli.conn.Send(input.Queue, contentType, input.Body, opts...)
	if err != nil {
		return err
	}

	return nil
}

func (cli *StompClient) makeSendOptions(input *pubsub.PublishInput) []func(*frame.Frame) error {
	opts := make([]func(*frame.Frame) error, 0)

	if input.Persistent {
		opts = append(opts, stomp3.SendOpt.Header("Persistent", "true"))
	}
	if input.MessageID != "" {
		opts = append(opts, stomp3.SendOpt.Header("message-id", input.MessageID))
	}

	return opts
}

// Subscribe subscribes to a queue.
func (cli *StompClient) Subscribe(input *pubsub.SubscribeInput) (pubsub.Subscription, error) {
	return cli.SubscribeWithContext(context.Background(), input)
}

// SubscribeWithContext subscribes to a queue with a context.
func (cli *StompClient) SubscribeWithContext(
	ctx context.Context,
	input *pubsub.SubscribeInput,
) (pubsub.Subscription, error) {
	cli.connMutex.RLock()
	defer cli.connMutex.RUnlock()

	if cli.conn == nil {
		if err := cli.Connect(); err != nil {
			return nil, err
		}
	}

	ack := stomp3.AckAuto
	if !input.AutoAck {
		ack = stomp3.AckClientIndividual
	}

	sub, err := cli.conn.Subscribe(input.Queue, ack)
	if err != nil {
		return nil, err
	}

	return cli.newSubscription(sub), nil
}

// Close closes the connection to the message broker.
func (cli *StompClient) Close() error {
	cli.connMutex.Lock()
	defer cli.connMutex.Unlock()

	if cli.conn == nil {
		return nil
	}

	defer func() {
		cli.conn = nil
	}()

	return cli.conn.Disconnect()
}
