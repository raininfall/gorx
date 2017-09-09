package observable

import (
	"time"

	"github.com/raininfall/gorx"
)

type observableInterval struct {
	d time.Duration
}

func (oba *observableInterval) Subscribe(out rx.InObserver) {
	go func() {
		defer out.Close()
		ticker := time.NewTicker(oba.d)
		defer ticker.Stop()
		for i := 0; true; i++ {
			select {
			case <-out.OnUnsubscribe():
				return
			case <-ticker.C:
				out.In() <- i
			}
		}
	}()
}

/*Interval will emit every duration time*/
func Interval(d time.Duration) rx.Observable {
	return &observableInterval{
		d: d,
	}
}
