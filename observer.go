package rx

/*InObserver input side of observer*/
type InObserver interface {
	Next() chan<- interface{}
	Complete()
	OnDispose() <-chan bool
}

/*OutObserver output side of observer*/
type OutObserver interface {
	OnNext() <-chan interface{}
	Dispose() chan<- bool
}

/*Observer is an interface for a consumer of push-based notifications delivered by an Observable.*/
type Observer interface {
	InObserver
	OutObserver
}
