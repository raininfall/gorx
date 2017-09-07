package subscription

import (
	"github.com/raininfall/gorx/teardown-logic"
)

/*Subscription represents a disposable resource, such as the execution of an Observable.*/
type Subscription interface {
	IsClosed() bool
	Add(teardownLogic.TeardownLogic) Subscription
	Remove(Subscription)
	Unsubscribe()
}

/*UnsubscribeFunc will be called when observer unsubscribe*/
type UnsubscribeFunc func()

type subscription struct {
	closed        bool //TODO: multi thread safety,volatile?
	teardownLogic teardownLogic.TeardownLogic
	tearDowns     []Subscription
}

/*New return instance of Subscription*/
func New(teardownLogic teardownLogic.TeardownLogic) Subscription {
	return &subscription{
		teardownLogic: teardownLogic,
	}
}

func (s *subscription) IsClosed() bool {
	return s.closed
}

func (s *subscription) Add(tl teardownLogic.TeardownLogic) Subscription {
	sub := New(tl)
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
	if s.closed {
		return /*Only cleanup once*/
	}
	if s.teardownLogic != nil {
		s.teardownLogic()
	}
	for _, td := range s.tearDowns {
		td.Unsubscribe()
	}
	s.closed = true
}
