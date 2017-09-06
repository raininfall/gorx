package observer

/*Observer is n interface for a consumer of push-based notifications delivered by an Observable.*/
type Observer interface {
	IsClosed() bool
	Next(interface{})
	Error(error)
	Complete()
}

/*IsClosedFunc Observer interface func*/
type IsClosedFunc func() bool

/*NextFunc Observer interface func*/
type NextFunc func(interface{})

/*ErrorFunc Observer interface func*/
type ErrorFunc func(error)

/*CompleteFunc Observer interface func*/
type CompleteFunc func()

type observer struct {
	closed   bool
	isClosed IsClosedFunc
	next     NextFunc
	err      ErrorFunc
	complete CompleteFunc
}

var defaultObserver = observer{
	isClosed: func() bool { return false },
	next:     func(interface{}) {},
	err:      func(error) {},
	complete: func() {},
}

func (ob *observer) IsClosed() bool {
	return ob.isClosed()
}

func (ob *observer) Next(value interface{}) {
	ob.next(value)
}

func (ob *observer) Error(err error) {
	ob.err(err)
}

func (ob *observer) Complete() {
	ob.complete()
}

/*New return default implement of Observer*/
func New(fns ...interface{}) Observer {
	ob := defaultObserver
	for _, fn := range fns {
		switch fn := fn.(type) {
		case IsClosedFunc:
			ob.isClosed = fn
		case NextFunc:
			ob.next = fn
		case ErrorFunc:
			ob.err = fn
		case CompleteFunc:
			ob.complete = fn
		}
	}
	return &ob
}