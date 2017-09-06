package rx

/*Same package*/

/*Observable is representation of any set of values over any amount of time*/
type Observable interface {
	Subscribe(Subscriber)
}

type observable struct {
	next <-chan interface{}
}

type observerToObservable struct {
	to     chan<- interface{}
	closed bool
}

func (observer *observerToObservable) IsClosed() bool {
	return observer.closed
}

func (observer *observerToObservable) Next(interface{}) {

}

func (observer *observerToObservable) Error(error) {

}

func (observer *observerToObservable) Complete() {

}

/*Create a new Observable, that will execute the specified function when an Observer subscribes to it*/
func Create(tl TeardownLogic) (Observable, Observer) {

	return nil, nil
}

func (observable *observable) Subscribe(subscriber Subscriber) {
	go func() {
		for item := range observable.next {
			if subscriber.IsClosed() {
				break
			}
			switch item := item.(type) {
			case error:
				subscriber.Error(item)
				return
			default:
				subscriber.Next(item)
			}
		}
		subscriber.Complete()
	}()
}
