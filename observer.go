package rx

/*Observer is n interface for a consumer of push-based notifications delivered by an Observable.*/
type Observer interface {
	IsClosed() bool
	Next(interface{})
	Error(error)
	Complete()
}

type observer struct {
	closed   bool
	isClosed ObserverIsClosedFunc
	next     ObserverNextFunc
	err      ObserverErrorFunc
	complete ObserverCompleteFunc
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

/*CreateObserver return default implement of Observer*/
func CreateObserver(fns ...interface{}) Observer {
	ob := defaultObserver
	for _, fn := range fns {
		switch fn := fn.(type) {
		case ObserverIsClosedFunc:
			ob.isClosed = fn
		case ObserverNextFunc:
			ob.next = fn
		case ObserverErrorFunc:
			ob.err = fn
		case ObserverCompleteFunc:
			ob.complete = fn
		}
	}
	return &ob
}
