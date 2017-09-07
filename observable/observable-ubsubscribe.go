package observable

type observableUnsubscribe chan int

func newObservableUnsubscribe() observableUnsubscribe {
	return make(chan int, 1)
}

func (ou observableUnsubscribe) notify() {
	ou <- 0
}

func (ou observableUnsubscribe) wait() <-chan int {
	return ou
}
