package observable

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/raininfall/gorx/observer"
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

	fin.Add(1)
	ob, obIn := New(nil)

	ob.Subscribe(obOut)

	obIn.Next(1)
	obIn.Next(2)
	obIn.Next(3)

	obIn.Complete()
	obIn.Error(errors.New("One"))
	fin.Wait()
	<-time.After(10 * time.Millisecond) /*wait observer all func called*/

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

	fin.Add(1)
	ob, obIn := New(nil)

	ob.Subscribe(obOut)

	obIn.Next(1)
	obIn.Next(2)
	obIn.Next(3)
	obIn.Error(errors.New("One"))

	fin.Wait()
	<-time.After(10 * time.Millisecond) /*wait observer all func called*/

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

	ob, obIn := New(nil)

	sub := ob.Subscribe(obOut)

	obIn.Next(1)
	obIn.Next(2)
	sub.Unsubscribe()
	fin := &sync.WaitGroup{}
	fin.Add(1)
	go func() {
		obIn.Next(3)
	}()
	<-time.After(10 * time.Millisecond) /*wait observer all func called*/

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
	sub := Interval(100 * time.Millisecond).Subscribe(obOut)
	<-time.After(450 * time.Millisecond)
	sub.Unsubscribe()
	<-time.After(150 * time.Millisecond)

	assert.Exactly(values, []int{0, 1, 2, 3})
	assert.Exactly(err.Error(), "None")
	assert.Exactly(done, "Not")
}

func TestObservableOf(t *testing.T) {
	assert := assert.New(t)
	fin := sync.WaitGroup{}
	out := make(chan interface{})
	obOut := observer.WrapChannel(out)
	values := []interface{}{}

	fin.Add(1)
	go func(c chan interface{}) {
		for item := range c {
			switch item := item.(type) {
			case error:
				t.Fatal("Should not output error")
			default:
				values = append(values, item)
			}
		}
		fin.Done()
	}(out)

	o1 := Of(1, 2, 3)
	o2 := Of(4, 5, 6)
	o3 := Of(7, 8, 9)

	Zip(o1, o2, o3).Subscribe(obOut)
	fin.Wait()

	assert.Exactly(values, []interface{}{
		[]interface{}{1, 4, 7},
		[]interface{}{2, 5, 8},
		[]interface{}{3, 6, 9},
	})

}