package observer

import "github.com/raininfall/gorx"

type observer struct {
	next    chan interface{}
	dispose chan bool
}

/*New will return standard impl Observer*/
func New(bufSize int) rx.Observer {
	return &observer{
		next:    make(chan interface{}, bufSize),
		dispose: make(chan bool, 1),
	}
}

func (o *observer) In() chan<- interface{} {
	return o.next
}

func (o *observer) Out() <-chan interface{} {
	return o.next
}

func (o *observer) Close() {
	close(o.next)
}

func (o *observer) Unsubscribe() {
	o.dispose <- true
}

func (o *observer) OnUnsubscribe() <-chan bool {
	return o.dispose
}
