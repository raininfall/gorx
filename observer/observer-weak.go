package observer

import rx "github.com/raininfall/gorx"

type observerWeak struct {
	next    chan interface{}
	dispose chan bool
}

/*NewWeak will return weak impl Observer*/
func NewWeak(bufSize int) rx.Observer {
	return &observer{
		next:    make(chan interface{}, bufSize),
		dispose: make(chan bool, 1),
	}
}

func (o *observerWeak) In() chan<- interface{} {
	return o.next
}

func (o *observerWeak) Out() <-chan interface{} {
	return o.next
}

func (o *observerWeak) Close() {
}

func (o *observerWeak) Unsubscribe() {
	close(o.next)
	close(o.dispose)
}

func (o *observerWeak) OnUnsubscribe() <-chan bool {
	return o.dispose
}
