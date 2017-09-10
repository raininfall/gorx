package observable

import (
	"errors"
	"testing"
	"time"

	"github.com/raininfall/gorx/observer"

	"github.com/stretchr/testify/assert"
)

func TestObservableMap(t *testing.T) {
	assert := assert.New(t)

	obs := observer.New(0)
	Of(1, 2, 3).Map(func(v interface{}) interface{} {
		return v.(int) * 2
	}).Subscribe(obs)

	values := []int{}
	for v := range obs.Out() {
		values = append(values, v.(int))
	}

	assert.Exactly([]int{2, 4, 6}, values)
}

func TestObservableMapError(t *testing.T) {
	assert := assert.New(t)

	obs := observer.New(0)
	Of(1, 2, 3, errors.New("Bang")).Map(func(v interface{}) interface{} {
		return v.(int) * 2
	}).Subscribe(obs)

	values := []int{}
	for v := range obs.Out() {
		switch v := v.(type) {
		case int:
			values = append(values, v)
		case error:
			assert.Exactly(errors.New("Bang"), v)
		}
	}

	assert.Exactly([]int{2, 4, 6}, values)
}

func TestObservableMapUnsubscribe(t *testing.T) {
	assert := assert.New(t)

	obs := observer.New(0)
	Interval(100 * time.Millisecond).Map(func(v interface{}) interface{} {
		return v.(int) * 2
	}).Subscribe(obs)

	go func() {
		<-time.After(350 * time.Millisecond)
		obs.Unsubscribe()
	}()

	values := []int{}
	for v := range obs.Out() {
		values = append(values, v.(int))
	}

	assert.Exactly([]int{0, 2, 4}, values)
}
