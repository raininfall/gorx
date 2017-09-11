package observable

import (
	"sync"

	"github.com/raininfall/gorx"
	"github.com/raininfall/gorx/observer"
)

type observableMergeMap struct {
	observable
	input *observable
	apply rx.MergeMapFunc
}

func (oba *observableMergeMap) Subscribe(out rx.InObserver) {
	outMerged := observer.New(0)
	observer.Pipe(outMerged, out)

	fin := &sync.WaitGroup{}

	fin.Add(1)
	go func() {
		defer fin.Done()

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
					newOba := oba.apply(item)
					newObs := observer.New(0)
					newOba.Subscribe(newObs)
					done := observer.WeakPipe(newObs, outMerged)
					fin.Add(1)
					go func() {
						<-done
						fin.Done()
					}()
				}
			case <-out.OnUnsubscribe():
				return
			}
		}

	}()

	go func() {
		fin.Wait()
		outMerged.Close()
	}()
}

func (oba *observable) MergeMap(apply rx.MergeMapFunc) rx.Observable {
	me := &observableMergeMap{
		input: oba,
		apply: apply,
	}
	me.subscribe = me.Subscribe

	return me
}
