package observer

type channelObserver struct {
	closed bool
	next   chan<- interface{}
}

/*WrapChannel will emit values to channel*/
func WrapChannel(c chan<- interface{}) Observer {
	return &channelObserver{
		next: c,
	}
}

func (observer *channelObserver) IsClosed() bool {
	return observer.closed
}

func (observer *channelObserver) Next(value interface{}) {
	if observer.closed {
		return
	}
	observer.next <- value
}

func (observer *channelObserver) Error(err error) {
	if observer.closed {
		return
	}
	observer.next <- err
	observer.close()
}

func (observer *channelObserver) Complete() {
	if observer.closed {
		return
	}
	observer.close()
}

/*close observable*/
func (observer *channelObserver) close() {
	close(observer.next)
	observer.closed = true
}