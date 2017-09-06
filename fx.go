package rx

type (
	/*ObserverIsClosedFunc Observer interface func*/
	ObserverIsClosedFunc func() bool
	/*ObserverNextFunc Observer interface func*/
	ObserverNextFunc func(interface{})
	/*ObserverErrorFunc Observer interface func*/
	ObserverErrorFunc func(error)
	/*ObserverCompleteFunc Observer interface func*/
	ObserverCompleteFunc func()
)
