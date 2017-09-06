package rx

/*Subscription represents a disposable resource, such as the execution of an Observable.*/
type Subscription interface {
	IsClosed() bool
	Add(TeardownLogic)
	Remove(Subscription)
	Unsubscribe()
}

type subscription struct {
	closed bool
}

/*CreateSubscription return instance of Subscription*/
func CreateSubscription() Subscription {
	return &subscription{}
}

func (s *subscription) IsClosed() bool {
	return s.closed
}

func (s *subscription) Add(TeardownLogic) {
	//TODO:
}

func (s *subscription) Remove(Subscription) {
	//TODO:
}

func (s *subscription) Unsubscribe() {
	//TODO:
}
