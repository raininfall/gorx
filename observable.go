package rx

/*Observable is a representation of any set of values over any amount of time*/
type Observable interface {
	Subscribe(InObserver)
	Map(MapFunc) Observable
	MergeMap(MergeMapFunc) Observable
}

/*MapFunc used by Map*/
type MapFunc func(interface{}) interface{}

/*MergeMapFunc used by MergeMap*/
type MergeMapFunc func(interface{}) Observable
