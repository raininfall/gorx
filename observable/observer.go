package observable

type observer struct {
	next    chan interface{}
	dispose chan bool
}

func (o *observer) Next() chan interface{} {
	return o.next
}

func (o *observer) Complete() {
	close(o.next)
}

func (o *observer) Dispose() chan bool {
	return o.dispose
}
