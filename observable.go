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
	closed bool
	next   chan<- interface{}
}

func (observer *observerToObservable) IsClosed() bool {
	return observer.closed
}

func (observer *observerToObservable) Next(value interface{}) {
	if observer.closed {
		return
	}
	observer.next <- value
}

func (observer *observerToObservable) Error(err error) {
	if observer.closed {
		return
	}
	observer.next <- err
	observer.close()
}

func (observer *observerToObservable) Complete() {
	if observer.closed {
		return
	}
	observer.close()
}

func (observer *observerToObservable) close() {
	close(observer.next)
	observer.closed = true
}

/*CreateObservable will create a new Observable, that will execute the specified function when an Observer subscribes to it*/
func CreateObservable(tl TeardownLogic) (Observable, Observer) {
	next := make(chan interface{})

	return &observable{
			next: next,
		}, &observerToObservable{
			closed: false,
			next:   next,
		}
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