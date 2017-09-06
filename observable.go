package rx

/*Same package*/

/*Observable is representation of any set of values over any amount of time*/
type Observable interface {
	Subscribe() Subscriber
}

type observable struct {
	next chan<- interface{}
}

/*Create a new Observable, that will execute the specified function when an Observer subscribes to it*/
func Create() Observable {
	next := make(chan interface{})
	return &observable{
		next: next,
	}
}

func (observable *observable) Subscribe() Subscriber {
	return nil
}
