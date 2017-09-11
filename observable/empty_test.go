package observable

import (
	"testing"

	"github.com/raininfall/gorx/observer"
	"github.com/stretchr/testify/assert"
)

func TestObservableEmpty(t *testing.T) {
	assert := assert.New(t)

	obs := observer.New(0)
	Empty().Subscribe(obs)

	count := 0
	for range obs.Out() {
		count++
	}

	assert.Equal(0, count)
}
