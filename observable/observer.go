package observable

type observer struct {
	next    chan interface{}
	dispose chan bool
}

func newObserver(bufSize int) *observer {
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
