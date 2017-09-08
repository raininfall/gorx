package observable

import (
	"github.com/eapache/channels"
	"github.com/raininfall/gorx"
)

type observable struct {
	in *observer
}

func (oba *observable) Subscribe(out rx.InObserver) {
	channels.Pipe(channels.Wrap(oba.in.Next()), channels.NativeInChannel(out.Next()))
}

/*Create a new Observable, that will execute the specified function when an Observer subscribes to it.*/
func Create(c chan<- Observer) rx.Observable {
	in := &observer{
		next:    make(chan interface{}),
		dispose: make(chan bool, 1),
	}
	c <- in

	return &observable{
		in: in,
	}
}
