package observable

import (
	"github.com/raininfall/gorx/observer"
	"github.com/raininfall/gorx/subscriber"
	"github.com/raininfall/gorx/teardown-logic"
)

type observableClean struct {
	o Observable
}

/*NewObservableClean will cleanup observable of limit value*/
func NewObservableClean(observable Observable) teardownLogic.TeardownLogic {
	o := &observableClean{
		o: observable,
	}

	return o.clean
}

func (oc *observableClean) clean() {
	ob := observer.New()
	sub := subscriber.New(ob, nil)
	oc.o.Subscribe(sub)
}
