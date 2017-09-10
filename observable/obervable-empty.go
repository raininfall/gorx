package observable

import (
	"github.com/raininfall/gorx"
)

type observableEmpty struct {
	observable
}

func (ob *observableEmpty) Subscribe(out rx.InObserver) {
	out.Close()
}

/*Empty will return a empty observable*/
func Empty() rx.Observable {
	me := &observableEmpty{}
	me.subscribe = me.Subscribe

	return me
}
