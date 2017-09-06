package rx

/*Subscription represents a disposable resource, such as the execution of an Observable.*/
type Subscription interface {
	IsClosed() bool
	Add(TeardownLogic) Subscription
	Remove(Subscription)
	Unsubscribe()
}

type subscription struct {
	closed      bool
	unsubscribe UnsubscribeFunc
	tearDowns   []Subscription
}

/*CreateSubscription return instance of Subscription*/
func CreateSubscription(unsubscribe UnsubscribeFunc) Subscription {
	return &subscription{
		unsubscribe: unsubscribe,
	}
}

func (s *subscription) IsClosed() bool {
	return s.closed
}

func (s *subscription) Add(tl TeardownLogic) Subscription {
	sub := CreateSubscription(UnsubscribeFunc(tl))
	s.tearDowns = append(s.tearDowns, sub)
	return sub
}

func (s *subscription) Remove(sub Subscription) {
	for i, td := range s.tearDowns {
		if td == sub {
			s.tearDowns = append(s.tearDowns[:i], s.tearDowns[i+1:]...)
		}
	}
}

func (s *subscription) Unsubscribe() {
	if s.unsubscribe != nil {
		s.unsubscribe()
	}
	for _, td := range s.tearDowns {
		td.Unsubscribe()
	}
	s.closed = true
}
