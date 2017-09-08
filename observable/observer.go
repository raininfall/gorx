package observable

type observer struct {
	closed  bool
	next    chan interface{}
	dispose chan bool
}

var unsubscribedChannel = make(chan interface{})
var errorChannel = make(chan interface{})

func init() {
	close(unsubscribedChannel)
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
	o.closed = true
	o.dispose <- true
}

func (o *observer) OnUnsubscribe() <-chan bool {
	return o.dispose
}
