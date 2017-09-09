package observable

import (
	"reflect"

	"github.com/raininfall/gorx"
	"github.com/raininfall/gorx/observer"
)

type observableMerge struct {
	inputs []rx.Observable
}

func (oba *observableMerge) Subscribe(out rx.InObserver) {
	go func() {
		defer out.Close()
		//The first elem holds for unsubscribe chan
		observers := make([]rx.Observer, len(oba.inputs)+1)
		cases := make([]reflect.SelectCase, len(oba.inputs)+1)

		for i, in := range oba.inputs {
			obs := observer.New(0)
			in.Subscribe(obs)
			defer obs.Unsubscribe()
			observers[i+1] = obs
			cases[i+1].Dir = reflect.SelectRecv
		}
		cases[0].Dir = reflect.SelectRecv

		for len(observers) > 1 {
			// Re-assign channel from func Out in case of change
			for i := 1; i < len(cases); i++ {
				cases[i].Chan = reflect.ValueOf(observers[i].Out())
			}
			cases[0].Chan = reflect.ValueOf(out.OnUnsubscribe())
			chosen, recv, ok := reflect.Select(cases)
			if !ok {
				//It should not be first elem(unsubscribe chan)
				observers = append(observers[:chosen], observers[chosen+1:]...)
				cases = append(cases[:chosen], cases[chosen+1:]...)
				continue
			}
			if chosen == 0 {
				return
			}
			out.In() <- recv.Interface()
			switch recv.Interface().(type) {
			case error:
				return
			}
		}
	}()
}

/*Merge all emitter into one observable*/
func Merge(inputs ...rx.Observable) rx.Observable {
	return &observableMerge{
		inputs: inputs,
	}
}
