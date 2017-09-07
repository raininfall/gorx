package observable

import (
	"time"

	"github.com/raininfall/gorx/observer"
	"github.com/raininfall/gorx/subscriber"
	"github.com/raininfall/gorx/teardown-logic"
)

/*Observable is representation of any set of values over any amount of time*/
type Observable interface {
	Subscribe(observer observer.Observer) subscriber.Subscriber
}

type observable struct {
	next <-chan interface{}
	tl   teardownLogic.TeardownLogic
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

/*close observable*/
func (observer *observerToObservable) close() {
	close(observer.next)
	observer.closed = true
}

/*New will create a new Observable, that will execute the specified function when an Observer subscribes to it*/
func New(tl teardownLogic.TeardownLogic) (Observable, observer.Observer) {
	next := make(chan interface{})

	return &observable{
			next: next,
			tl:   tl,
		}, &observerToObservable{
			closed: false,
			next:   next,
		}

}

func (observable *observable) Subscribe(observer observer.Observer) subscriber.Subscriber {
	ou := newObservableUnsubscribe()
	subscriber := subscriber.New(observer, ou.notify)
	if observable.tl != nil {
		subscriber.Add(observable.tl)
	}

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

	return subscriber
}

/*Interval will emit every \a d time, unavailable when first unsubscribe*/
func Interval(d time.Duration) Observable {
	ou := newObservableUnsubscribe()
	observable, observer := New(ou.notify)
	go func() {
		ticker := time.NewTicker(d)
		defer ticker.Stop()
		for i := 0; !observer.IsClosed(); {
			select {
			case <-ou.wait(): /*There will bo no complete or error emit from this observable*/
				return
			case <-ticker.C:
				observer.Next(i)
				i++
			}
		}
	}()

	return observable
}

/*Of will emit all values in seq*/
func Of(items ...interface{}) Observable {
	observable, observer := New(nil)
	go func() {
		for _, item := range items {
			observer.Next(item)
		}
		observer.Complete()
	}()

	return observable
}

/*Zip will combines multiple Observables to create an Observable whose values are calculated from the values, in order, of each of its input Observables.*/
func Zip(observables ...observable) Observable {
	return nil
}
