package observable

import (
	"github.com/raininfall/channels"
	"github.com/raininfall/gorx"
)

type observable struct {
	in *observer
}

func (oba *observable) Subscribe(out rx.Observer) {
	channels.Pipe(oba.in, out)
}

/*Create a new Observable, that will execute the specified function when an Observer subscribes to it.*/
func Create(c chan<- rx.Observer) rx.Observable {
	in := newObserver(10)
	c <- in

	return &observable{
		in: in,
	}
}
