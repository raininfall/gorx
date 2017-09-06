package rx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObserver(t *testing.T) {
	assert := assert.New(t)
	nums := []int{}
	var errGot error
	errSet := errors.New("Done")
	done := "Not"

	nextFn := ObserverNextFunc(func(value interface{}) {
		nums = append(nums, value.(int))
	})
	errorFn := ObserverErrorFunc(func(err error) {
		errGot = err
	})
	doneFn := ObserverCompleteFunc(func() {
		done = "Done"
	})

	/*Mocking observable ops*/
	ob := CreateObserver(nextFn, errorFn, doneFn)
	ob.Next(1)
	ob.Next(2)
	ob.Next(3)
	ob.Error(errSet)
	ob.Complete()

	assert.Exactly([]int{1, 2, 3}, nums)
	assert.Exactly(errGot, errSet)
	assert.Exactly(done, "Done")
}
