package observable

import (
	"testing"

	"github.com/raininfall/gorx/observer"
	"github.com/stretchr/testify/assert"
)

func TestObservableZip(t *testing.T) {
	assert := assert.New(t)

	obs := observer.New(0)
	Zip(
		Of(1, 2, 3),
		Of(4, 5, 6),
		Of(7, 8, 9),
	).Subscribe(obs)

	values := []interface{}{}
	for v := range obs.Out() {
		values = append(values, v)
	}

	assert.Exactly([]interface{}{
		[]interface{}{1, 4, 7},
		[]interface{}{2, 5, 8},
		[]interface{}{3, 6, 9},
	}, values)
}
