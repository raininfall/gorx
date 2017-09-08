package rx

/*Observable is a representation of any set of values over any amount of time*/
type Observable interface {
	Subscribe(InObserver)
}
