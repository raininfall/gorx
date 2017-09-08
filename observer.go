package rx

/*InObserver input side of observer*/
type InObserver interface {
	In() chan<- interface{}
	Close()
	OnUnsubscribe() <-chan bool
}

/*OutObserver output side of observer*/
type OutObserver interface {
	Out() <-chan interface{}
	Unsubscribe()
}

/*Observer is an interface for a consumer of push-based notifications delivered by an Observable.*/
type Observer interface {
	InObserver
	OutObserver
}
