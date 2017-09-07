package subscriber

import (
	"github.com/raininfall/gorx/observer"
	"github.com/raininfall/gorx/subscription"
	"github.com/raininfall/gorx/teardown-logic"
)

/*Subscriber Implements the Observer interface and the Subscription interface*/
type Subscriber interface {
	Add(teardownLogic.TeardownLogic)
	Remove(subscription.Subscription)
	Unsubscribe()
	observer.Observer
}

type subscriber struct {
	subscription subscription.Subscription
	observer     observer.Observer
}

/*New return instance Subscriber*/
func New(observer observer.Observer) Subscriber {
	return &subscriber{
		subscription: subscription.New(nil),
		observer:     observer,
	}
}

func (s *subscriber) IsClosed() bool {
	return s.subscription.IsClosed()
}

func (s *subscriber) Add(tl teardownLogic.TeardownLogic) {
	s.subscription.Add(tl)
}

func (s *subscriber) Remove(sub subscription.Subscription) {
	s.subscription.Remove(sub)
}

func (s *subscriber) Unsubscribe() {
	s.subscription.Unsubscribe()
}

func (s *subscriber) Next(value interface{}) {
	s.observer.Next(value)
}
func (s *subscriber) Error(err error) {
	s.observer.Error(err)
}

func (s *subscriber) Complete() {
	s.observer.Complete()
}
