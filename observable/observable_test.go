package observable

import (
	"runtime"
	"sync"
	"testing"

	"github.com/raininfall/gorx"
	"github.com/raininfall/gorx/observer"
	"github.com/stretchr/testify/assert"
)

func TestObservableCreate(t *testing.T) {
	assert := assert.New(t)
	fin := &sync.WaitGroup{}
	numGoroutines := runtime.NumGoroutine()

	cb := make(chan rx.InObserver, 1) /*Avoid observable.Create block*/
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
		obs := observer.New(0)
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

func TestObserverUnsubscribe(t *testing.T) {
	assert := assert.New(t)
	fin := &sync.WaitGroup{}
	numGoroutines := runtime.NumGoroutine()

	cb := make(chan rx.InObserver, 1) /*Avoid observable.Create block*/
	oba := Create(cb)

	fin.Add(1)
	go func() {
		defer fin.Done()
		obs := <-cb
		defer obs.Close()
		for i := 1; i > 0; i++ {
			select {
			case <-obs.OnUnsubscribe():
				return
			case obs.In() <- i:
			}
		}
	}()

	values := []int{}
	fin.Add(1)
	go func() {
		defer fin.Done()
		obs := observer.New(0)
		oba.Subscribe(obs)
		for item := range obs.Out() {
			values = append(values, item.(int))
			if len(values) == 3 {
				obs.Unsubscribe()
				break
			}
		}
		/*It should be OK not to clean the observer.Out() after Unsubscribe*/
		assert.Exactly([]int{1, 2, 3}, values)
	}()

	fin.Wait()
	assert.Exactly(runtime.NumGoroutine(), numGoroutines)
}
