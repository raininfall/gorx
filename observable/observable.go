package observable

import (
	"github.com/raininfall/gorx"
)

type observable struct {
	onSubscribe chan<- rx.InObserver
}

func (oba *observable) Subscribe(out rx.InObserver) {
	oba.onSubscribe <- out
}

/*Create a new Observable, that will execute the specified function when an Observer subscribes to it.*/
func Create(onSubscribe chan<- rx.InObserver) rx.Observable {
	return &observable{
		onSubscribe: onSubscribe,
	}
}
