package rx

/*InObserver input side of observer*/
type InObserver interface {
	In() chan<- interface{}
	Close()
	OnDispose() <-chan bool
}

/*OutObserver output side of observer*/
type OutObserver interface {
	Out() <-chan interface{}
	Dispose() chan<- bool
}

/*Observer is an interface for a consumer of push-based notifications delivered by an Observable.*/
type Observer interface {
	InObserver
	OutObserver
}
