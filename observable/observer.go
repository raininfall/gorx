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

func (o *observer) Next() chan<- interface{} {
	return o.next
}

func (o *observer) OnNext() <-chan interface{} {
	return o.next
}

func (o *observer) Complete() {
	close(o.next)
}

func (o *observer) Dispose() chan<- bool {
	return o.dispose
}

func (o *observer) OnDispose() <-chan bool {
	return o.dispose
}
