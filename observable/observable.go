package observable

import rx "github.com/raininfall/gorx"

type observable struct {
	subscribe func(rx.InObserver)
}
