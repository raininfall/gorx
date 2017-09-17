package observable

import (
	rx "github.com/raininfall/gorx"
	"github.com/raininfall/gorx/observer"
)

type observableZip struct {
	observable
	values []rx.Observable
}

func (me *observableZip) Subscribe(out rx.InObserver) {
	go func() {
		defer out.Close()

		observers := make([]rx.Observer, len(me.values))
		for i := 0; i < len(observers); i++ {
			observers[i] = observer.New(0)
			me.values[i].Subscribe(observers[i])
		}
		for {
			zip := make([]interface{}, len(me.values))
			for i, observer := range observers {
				select {
				case <-out.OnUnsubscribe():
					return
				case item, ok := <-observer.Out():
					if !ok {
						return
					}
					switch item := item.(type) {
					case error:
						return
					default:
						zip[i] = item
					}
				}
			}
			out.In() <- zip
		}
	}()
}

//Zip will merge every emit from observable
func Zip(observables ...rx.Observable) rx.Observable {
	me := &observableZip{
		values: observables,
	}
	me.subscribe = me.Subscribe

	return me
}
