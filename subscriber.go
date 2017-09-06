package rx

/*Subscriber Implements the Observer interface and the Subscription interface*/
type Subscriber interface {
	Add(TeardownLogic)
	Remove(Subscription)
	Unsubscribe()
	Observer
}

type subscriber struct {
	subscription Subscription
	observer     Observer
}

func (s *subscriber) IsClosed() bool {
	return s.subscription.IsClosed()
}

func (s *subscriber) Add(tl TeardownLogic) {
	s.subscription.Add(tl)
}

func (s *subscriber) Remove(sub Subscription) {
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
