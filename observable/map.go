package observable

import (
	"github.com/raininfall/gorx"
	"github.com/raininfall/gorx/observer"
)

type observableMap struct {
	observable
	input *observable
	apply rx.MapFunc
}

func (oba *observableMap) Subscribe(out rx.InObserver) {
	go func() {
		defer out.Close()
		input := observer.New(0)
		oba.input.subscribe(input)

		for {
			select {
			case item, ok := <-input.Out():
				if !ok {
					return
				}
				switch item.(type) {
				case error:
					out.In() <- item
					return
				default:
					out.In() <- oba.apply(item)
				}
			case <-out.OnUnsubscribe():
				return
			}
		}

	}()
}

func (oba *observable) Map(apply rx.MapFunc) rx.Observable {
	me := &observableMap{
		input: oba,
		apply: apply,
	}
	me.subscribe = me.Subscribe

	return me
}
