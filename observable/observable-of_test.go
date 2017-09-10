package observable

import (
	"errors"
	"testing"
	"time"

	"github.com/raininfall/gorx/observer"
	"github.com/stretchr/testify/assert"
)

func TestObservableOf(t *testing.T) {
	assert := assert.New(t)

	oba := Of(1, 2, 3, errors.New("Bang"), 4, nil)
	obs := observer.New(0)
	oba.Subscribe(obs)

	values := []int{}
	for v := range obs.Out() {
		switch v := v.(type) {
		case error:
			assert.Exactly(errors.New("Bang"), v)
			continue
		default:
			values = append(values, v.(int))
		}
	}

	assert.Exactly([]int{1, 2, 3}, values)
}

func TestObservableOfUnsubscribe(t *testing.T) {
	assert := assert.New(t)

	oba := Of(1, 2, 3, 4, 5, errors.New("Bang"), 6, nil)
	obs := observer.New(0)
	oba.Subscribe(obs)

	values := []int{}
	i := 0
	for v := range obs.Out() {
		switch v := v.(type) {
		case error:
			assert.Exactly(errors.New("Bang"), v)
			continue
		default:
			values = append(values, v.(int))
		}
		i++
		if i == 3 {
			<-time.After(50 * time.Millisecond) /*Give it some time to fill next value*/
			obs.Unsubscribe()
		}
	}

	assert.Exactly([]int{1, 2, 3}, values)
}
