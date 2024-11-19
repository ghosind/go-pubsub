package stomp

import (
	"github.com/ghosind/go-pubsub"
	stomp3 "github.com/go-stomp/stomp/v3"
)

// StompSubscription is a subscription for STOMP protocol.
type StompSubscription struct {
	// cli is the pub-sub client for STOMP protocol.
	cli *StompClient

	// subscription is the subscription for STOMP protocol.
	subscription *stomp3.Subscription

	// msgChan is the channel to receive messages.
	msgChan chan pubsub.Message

	// closeChan is the channel to notify the subscription closed.
	closeChan chan struct{}
}

// newSubscription creates a new subscription for STOMP protocol.
func (cli *StompClient) newSubscription(sub *stomp3.Subscription) *StompSubscription {
	subscription := new(StompSubscription)

	subscription.cli = cli
	subscription.subscription = sub

	go subscription.runLoop()

	return subscription
}

// Receive returns a channel that will receive messages from the subscription.
func (sub *StompSubscription) Receive() <-chan pubsub.Message {
	return sub.msgChan
}

// Unsubscribe unsubscribes from the subscription.
func (sub *StompSubscription) Unsubscribe() error {
	sub.closeChan <- struct{}{}
	return sub.subscription.Unsubscribe()
}

// runLoop runs the loop to receive messages or close notification of the subscription.
func (sub *StompSubscription) runLoop() {
	for {
		select {
		case msg := <-sub.subscription.C:
			if msg == nil {
				continue
			}
			if msg.Err != nil {
				if !sub.subscription.Active() {
					sub.reconnect()
				}
			} else {
				sub.msgChan <- newStompMessage(msg)
			}
		case <-sub.closeChan:
			return
		case <-sub.cli.closeChan:
			return
		}
	}
}

func (sub *StompSubscription) reconnect() {
	destination := sub.subscription.Destination()
	ackMode := sub.subscription.AckMode()

	if err := sub.cli.Connect(); err != nil {
		panic(err)
	}

	subscription, err := sub.cli.conn.Subscribe(destination, ackMode)
	if err != nil {
		panic(err)
	}
	sub.subscription = subscription
}
