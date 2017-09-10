package observable

import (
	"github.com/raininfall/gorx"
)

type observableCreate struct {
	observable
	onSubscribe chan<- rx.InObserver
}

func (oba *observableCreate) Subscribe(out rx.InObserver) {
	oba.onSubscribe <- out
}

/*Create a new Observable, that will execute the specified function when an Observer subscribes to it.*/
func Create(onSubscribe chan<- rx.InObserver) rx.Observable {
	me := &observableCreate{
		onSubscribe: onSubscribe,
	}
	me.subscribe = me.Subscribe

	return me
}
