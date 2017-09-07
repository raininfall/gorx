package observable

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/raininfall/gorx/observer"
	"github.com/raininfall/gorx/subscriber"
	"github.com/stretchr/testify/assert"
)

func TestObservableComplete(t *testing.T) {
	assert := assert.New(t)
	done := "Not"
	err := errors.New("None")
	values := []int{}
	fin := &sync.WaitGroup{}

	nextFn := observer.NextFunc(func(value interface{}) {
		values = append(values, value.(int))
	})
	doneFn := observer.CompleteFunc(func() {
		done = "Yes"
		fin.Done()
	})
	errFn := observer.ErrorFunc(func(errOut error) {
		err = errOut
	})

	obOut := observer.New(nextFn, doneFn, errFn)
	sub := subscriber.New(obOut)

	fin.Add(1)
	ob, obIn := New(nil)

	assert.Nil(ob.Subscribe(sub))

	obIn.Next(1)
	obIn.Next(2)
	obIn.Next(3)

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

	nextFn := observer.NextFunc(func(value interface{}) {
		values = append(values, value.(int))
	})
	errFn := observer.ErrorFunc(func(errOut error) {
		err = errOut
		fin.Done()
	})
	doneFn := observer.CompleteFunc(func() {
		done = "Yes"
		fin.Done()
	})

	obOut := observer.New(nextFn, errFn, doneFn)
	sub := subscriber.New(obOut)

	fin.Add(1)
	ob, obIn := New(nil)

	assert.Nil(ob.Subscribe(sub))

	obIn.Next(1)
	obIn.Next(2)
	obIn.Next(3)
	obIn.Error(errors.New("One"))

	fin.Wait()

	assert.Exactly(values, []int{1, 2, 3})
	assert.Exactly(err.Error(), "One")
	assert.Exactly(done, "Not")
	assert.Exactly(obIn.IsClosed(), true)
}

func TestObservableUnsubscribe(t *testing.T) {
	assert := assert.New(t)
	err := errors.New("None")
	done := "Not"
	values := []int{}
	fin := &sync.WaitGroup{}

	nextFn := observer.NextFunc(func(value interface{}) {
		values = append(values, value.(int))
	})
	errFn := observer.ErrorFunc(func(errOut error) {
		err = errOut
		t.Fatal("Should not call error")
	})
	doneFn := observer.CompleteFunc(func() {
		done = "Yes"
		t.Fatal("Should not call complete")
	})

	obOut := observer.New(nextFn, errFn, doneFn)
	sub := subscriber.New(obOut)

	ob, obIn := New(nil)

	assert.Nil(ob.Subscribe(sub))

	fin.Add(1)
	go func() {
		obIn.Next(1)
		obIn.Next(2)
		<-time.After(10 * time.Millisecond) /*wait next func called*/
		sub.Unsubscribe()
		fin.Done()
	}()

	<-time.After(500 * time.Millisecond) /*wait observer all func called*/
	fin.Wait()

	assert.Exactly(values, []int{1, 2})
	assert.Exactly(err.Error(), "None")
	assert.Exactly(done, "Not")
	assert.Exactly(obIn.IsClosed(), false)
}

func TestObservableInterval(t *testing.T) {
	assert := assert.New(t)
	err := errors.New("None")
	done := "Not"
	values := []int{}

	nextFn := observer.NextFunc(func(value interface{}) {
		values = append(values, value.(int))
	})
	errFn := observer.ErrorFunc(func(errOut error) {
		err = errOut
	})
	doneFn := observer.CompleteFunc(func() {
		done = "Yes"
	})

	obOut := observer.New(nextFn, errFn, doneFn)
	sub := subscriber.New(obOut)
	Interval(100 * time.Millisecond).Subscribe(sub)
	<-time.After(450 * time.Millisecond)

	assert.Exactly(values, []int{0, 1, 2, 3})
	assert.Exactly(err.Error(), "None")
	assert.Exactly(done, "Not")
}
