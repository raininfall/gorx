package observable

import (
	"time"

	"github.com/raininfall/gorx/observer"
	"github.com/raininfall/gorx/subscriber"
	"github.com/raininfall/gorx/teardown-logic"
)

/*Observable is representation of any set of values over any amount of time*/
type Observable interface {
	Subscribe(subscriber.Subscriber) error
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

/*New will create a new Observable, that will execute the specified function when an Observer subscribes to it*/
func New(tl teardownLogic.TeardownLogic) (Observable, observer.Observer) {
	next := make(chan interface{})

	return &observable{
			next: next,
		}, &observerToObservable{
			closed: false,
			next:   next,
		}
}

func (observable *observable) Subscribe(subscriber subscriber.Subscriber) error {
	/*Subscriber is protected for multi-thread outside this func ,it will work only if subscriber called here, not in go routine*/
	ou := newObservableUnsubscribe()
	subscriber.Add(ou.notify)

	go func() {
		for !subscriber.IsClosed() {
			select {
			case <-ou.wait():
				return /*Stop because of unsubscribe*/
			case item, notComplete := <-observable.next:
				if !notComplete {
					subscriber.Complete()
					return /*Stop because of complete*/
				}
				switch item := item.(type) {
				case error:
					subscriber.Error(item)
					return /*Stop because of error*/
				default:
					subscriber.Next(item)
				}
			}
		}
	}()

	return nil
}

/*Interval will emit every \a d time*/
func Interval(d time.Duration) Observable {
	observable, observer := New(nil)
	go func() {
		ticker := time.NewTicker(d)
		for i := 0; !observer.IsClosed(); i++ {
			<-ticker.C
			observer.Next(i)
		}
	}()
	return observable
}

/*Zip will combines multiple Observables to create an Observable whose values are calculated from the values, in order, of each of its input Observables.*/
func Zip(observables ...observable) Observable {
	return nil
}
