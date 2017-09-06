package rx

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObservableComplete(t *testing.T) {
	assert := assert.New(t)
	done := "Not"
	err := errors.New("None")
	values := []int{}
	fin := &sync.WaitGroup{}

	nextFn := ObserverNextFunc(func(value interface{}) {
		values = append(values, value.(int))
	})
	doneFn := ObserverCompleteFunc(func() {
		done = "Yes"
		fin.Done()
	})
	errFn := ObserverErrorFunc(func(errOut error) {
		err = errOut
	})

	obOut := CreateObserver(nextFn, doneFn, errFn)
	sub := CreateSubscriber(obOut)

	ob, obIn := CreateObservable(nil)

	ob.Subscribe(sub)

	obIn.Next(1)
	obIn.Next(2)
	obIn.Next(3)
	fin.Add(1)
	obIn.Complete()
	obIn.Error(errors.New("One"))
	fin.Wait()

	assert.Exactly(values, []int{1, 2, 3})
	assert.Exactly(done, "Yes")
	assert.Exactly(err.Error(), "None")
	assert.Exactly(obIn.IsClosed(), true)
}

func TestObservableError(t *testing.T) {
	assert := assert.New(t)
	err := errors.New("None")
	done := "Not"
	values := []int{}
	fin := &sync.WaitGroup{}

	nextFn := ObserverNextFunc(func(value interface{}) {
		values = append(values, value.(int))
	})
	errFn := ObserverErrorFunc(func(errOut error) {
		err = errOut
		fin.Done()
	})
	doneFn := ObserverCompleteFunc(func() {
		done = "Yes"
		fin.Done()
	})

	obOut := CreateObserver(nextFn, errFn, doneFn)
	sub := CreateSubscriber(obOut)

	ob, obIn := CreateObservable(nil)

	ob.Subscribe(sub)

	obIn.Next(1)
	obIn.Next(2)
	obIn.Next(3)
	fin.Add(1)
	obIn.Error(errors.New("One"))
	fin.Wait()

	assert.Exactly(values, []int{1, 2, 3})
	assert.Exactly(err.Error(), "One")
	assert.Exactly(done, "Not")
	assert.Exactly(obIn.IsClosed(), true)
}