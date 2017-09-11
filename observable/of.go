package observable

import (
	"github.com/raininfall/gorx"
)

type observableOf struct {
	observable
	values []interface{}
}

func (oba *observableOf) Subscribe(out rx.InObserver) {
	go func() {
		defer out.Close()
		for _, item := range oba.values {
			select {
			case <-out.OnUnsubscribe():
				return
			case out.In() <- item:
				switch item.(type) {
				case error:
					return
				}
			}
		}
	}()
}

/*Of a new Observable, that will execute the specified function when an Observer subscribes to it.*/
func Of(values ...interface{}) rx.Observable {
	me := &observableOf{
		values: values,
	}
	me.subscribe = me.Subscribe

	return me
}
