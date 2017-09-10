package observer

import "github.com/raininfall/gorx"

/*Pipe observers*/
func Pipe(in rx.OutObserver, out rx.InObserver) <-chan bool {
	done := make(chan bool)

	go func() {
		defer out.Close()
		weakPipeImpl(in, out, done)
	}()

	return done
}

/*WeakPipe observers without close*/
func WeakPipe(in rx.OutObserver, out rx.InObserver) <-chan bool {
	done := make(chan bool)

	go weakPipeImpl(in, out, done)

	return done
}

func weakPipeImpl(in rx.OutObserver, out rx.InObserver, done chan<- bool) {
	defer close(done)
	for {
		select {
		case <-out.OnUnsubscribe():
			in.Unsubscribe()
			return
		case item, ok := <-in.Out():
			if !ok {
				return
			}
			out.In() <- item
			switch item.(type) {
			case error:
				return
			}
		}
	}
}
