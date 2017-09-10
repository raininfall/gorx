package rx

/*Observable is a representation of any set of values over any amount of time*/
type Observable interface {
	Subscribe(InObserver)
	Map(MapFunc) Observable
}

/*MapFunc used by Map*/
type MapFunc func(interface{}) interface{}
