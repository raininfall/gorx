package observable

import (
	"github.com/eapache/channels"
	"github.com/raininfall/gorx"
)

type observable struct {
	in *observer
}

func (oba *observable) Subscribe(out rx.InObserver) {
	channels.Pipe(oba.in, out)
}

/*Create a new Observable, that will execute the specified function when an Observer subscribes to it.*/
func Create(c chan<- rx.Observer) rx.Observable {
	in := newObserver()
	c <- in

	return &observable{
		in: in,
	}
}
