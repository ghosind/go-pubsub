package stomp

import (
	"context"
	"errors"
	"io"
	"net"
	"sync"

	"github.com/ghosind/go-pubsub"
	stomp3 "github.com/go-stomp/stomp/v3"
)

// StompClient is a Pub-Sub client for STOMP protocol.
type StompClient struct {
	// conn is the STOMP connection.
	conn *stomp3.Conn

	netConn io.ReadWriteCloser

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

	if err := cli.establishNetConnection(); err != nil {
		return err
	}

	conn, err := stomp3.Connect(
		cli.netConn,
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
	return errors.New("not implemented")
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
	return nil, errors.New("not implemented")
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

func (cli *StompClient) establishNetConnection() error {
	if cli.netConn != nil {
		return nil
	}

	conn, err := net.Dial("tcp", cli.address)
	if err != nil {
		return err
	}

	cli.netConn = conn

	return nil
}
