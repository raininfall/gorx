package observable

import (
	"testing"
	"time"

	"github.com/raininfall/gorx/observer"

	"github.com/stretchr/testify/assert"
)

func TestObservableInterval(t *testing.T) {
	assert := assert.New(t)

	oba := Interval(100 * time.Millisecond)
	obs := observer.New(0)
	oba.Subscribe(obs)

	go func() {
		<-time.After(450 * time.Millisecond)
		obs.Unsubscribe()
	}()

	values := []int{}
	for v := range obs.Out() {
		values = append(values, v.(int))
	}

	assert.Exactly([]int{0, 1, 2, 3}, values)
}
