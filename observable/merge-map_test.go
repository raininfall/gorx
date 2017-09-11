package observable

import (
	"sort"
	"testing"

	"github.com/raininfall/gorx"

	"github.com/raininfall/gorx/observer"
	"github.com/stretchr/testify/assert"
)

func TestObservableMergeMap(t *testing.T) {
	assert := assert.New(t)

	obs := observer.New(0)
	Of(1, 2, 3).MergeMap(func(v interface{}) rx.Observable {
		return Of(v.(int) * 2)
	}).Subscribe(obs)

	values := []int{}
	for v := range obs.Out() {
		values = append(values, v.(int))
	}

	sort.Ints(values)
	assert.Exactly([]int{2, 4, 6}, values)
}
