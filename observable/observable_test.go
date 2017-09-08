package observable

import (
	"testing"

	"github.com/raininfall/gorx"
	"github.com/stretchr/testify/assert"
)

func TestObservableCreate(t *testing.T) {
	assert = assert.New(t)

	cb := make(chan rx.Observer)

	oba := Create(cb)

	go func() {
		obs := <-cb
		obs.Next() <- 1
		obs.Next() <- 2
		obs.Next() <- 3
		obs.Complete()
	}()
}
