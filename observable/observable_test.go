package observable

import (
	"runtime"
	"sync"
	"testing"

	"github.com/raininfall/gorx"
	"github.com/stretchr/testify/assert"
)

func TestObservableCreate(t *testing.T) {
	assert := assert.New(t)
	fin := &sync.WaitGroup{}
	numGoroutines := runtime.NumGoroutine()

	cb := make(chan rx.Observer, 1) /*Avoid observable.Create block*/
	oba := Create(cb)

	fin.Add(1)
	go func() {
		obs := <-cb
		obs.In() <- 1
		obs.In() <- 2
		obs.In() <- 3
		obs.Close()
		fin.Done()
	}()

	values := []int{}
	fin.Add(1)
	go func() {
		obs := newObserver()
		oba.Subscribe(obs)
		for item := range obs.Out() {
			values = append(values, item.(int))
		}
		fin.Done()
	}()

	fin.Wait()
	assert.Exactly([]int{1, 2, 3}, values)
	assert.Exactly(runtime.NumGoroutine(), numGoroutines)
}
