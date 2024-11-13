package stomp

import (
	"github.com/ghosind/go-pubsub"
	stomp3 "github.com/go-stomp/stomp/v3"
)

type StompSubscription struct {
	cli *StompClient

	subscription *stomp3.Subscription

	msgChan chan pubsub.Message

	closeChan chan struct{}
}

func (cli *StompClient) newSubscription(sub *stomp3.Subscription) *StompSubscription {
	subscription := new(StompSubscription)

	subscription.cli = cli
	subscription.subscription = sub

	go subscription.runLoop()

	return subscription
}

func (sub *StompSubscription) Receive() <-chan pubsub.Message {
	return sub.msgChan
}

func (sub *StompSubscription) Unsubscribe() error {
	sub.closeChan <- struct{}{}
	return sub.subscription.Unsubscribe()
}

func (sub *StompSubscription) runLoop() {
	for {
		select {
		case msg := <-sub.subscription.C:
			if msg == nil {
				continue
			}

			sub.msgChan <- newStompMessage(msg)
		case <-sub.closeChan:
			return
		case <-sub.cli.closeChan:
			return
		}
	}
}
