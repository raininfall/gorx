package observable

import (
	"errors"
	"sort"
	"testing"
	"time"

	"github.com/raininfall/gorx/observer"
	"github.com/stretchr/testify/assert"
)

func TestObservableMerge(t *testing.T) {
	assert := assert.New(t)

	oba1 := Of(1, 2, 3)
	oba2 := Of(4, 5, 6)
	obs := observer.New(0)
	Merge(oba1, oba2).Subscribe(obs)

	values := []int{}
	for v := range obs.Out() {
		values = append(values, v.(int))
	}

	sort.Ints(values)
	assert.Exactly([]int{1, 2, 3, 4, 5, 6}, values)
}

func TestObservableMergeUnsubscribe(t *testing.T) {
	assert := assert.New(t)

	oba1 := Of(1, 2, 3)
	oba2 := Interval(100 * time.Millisecond)
	obs := observer.New(0)
	Merge(oba1, oba2).Subscribe(obs)

	go func() {
		<-time.After(150 * time.Millisecond)
		obs.Unsubscribe()
	}()

	values := []int{}
	for v := range obs.Out() {
		values = append(values, v.(int))
	}

	assert.Exactly([]int{1, 2, 3, 0}, values)
}

func TestObservableMergeError(t *testing.T) {
	assert := assert.New(t)

	oba1 := Of(1, 2, 3, errors.New("Bang"))
	oba2 := Interval(100 * time.Millisecond)
	obs := observer.New(0)
	Merge(oba1, oba2).Subscribe(obs)

	values := []interface{}{}
	for v := range obs.Out() {
		values = append(values, v)
	}

	assert.Exactly([]interface{}{1, 2, 3, errors.New("Bang")}, values)
}
