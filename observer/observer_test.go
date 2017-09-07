package observer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObserver(t *testing.T) {
	assert := assert.New(t)
	all := []int{}
	var errGot error
	errSet := errors.New("Done")
	done := "Not"

	nextFn := NextFunc(func(value interface{}) {
		all = append(all, value.(int))
	})
	errorFn := ErrorFunc(func(err error) {
		errGot = err
	})
	doneFn := CompleteFunc(func() {
		done = "Done"
	})

	/*Mocking observable ops*/
	ob := New(nextFn, errorFn, doneFn)
	ob.Next(1)
	ob.Next(2)
	ob.Next(3)
	ob.Error(errSet)
	ob.Complete()

	assert.Exactly([]int{1, 2, 3}, all)
	assert.Exactly(errGot, errSet)
	assert.Exactly(done, "Done")
}