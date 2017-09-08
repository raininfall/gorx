package observable

type observer struct {
	next    chan interface{}
	dispose chan bool
}

func newObserver() *observer {
	return &observer{
		next:    make(chan interface{}),
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

func (o *observer) Dispose() chan<- bool {
	return o.dispose
}

func (o *observer) OnDispose() <-chan bool {
	return o.dispose
}
