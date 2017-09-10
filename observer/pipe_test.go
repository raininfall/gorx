package observer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipe(t *testing.T) {
	assert := assert.New(t)

	obs1 := New(0)
	obs2 := New(0)

	values := []int{}
	go func() {
		for v := range obs2.Out() {
			values = append(values, v.(int))
		}
	}()

	go func() {
		obs1.In() <- 1
		obs1.In() <- 2
		obs1.In() <- 3
		obs1.Close()
	}()

	<-Pipe(obs1, obs2)

	assert.Exactly([]int{1, 2, 3}, values)
}

func TestPipeError(t *testing.T) {
	assert := assert.New(t)

	obs1 := New(0)
	obs2 := New(0)
	err := errors.New("Not")

	values := []int{}
	go func() {
		for v := range obs2.Out() {
			switch v := v.(type) {
			case error:
				err = v
			default:
				values = append(values, v.(int))
			}
		}
	}()

	go func() {
		obs1.In() <- 1
		obs1.In() <- 2
		obs1.In() <- 3
		obs1.In() <- errors.New("Bang")
		obs1.Close()
	}()

	<-Pipe(obs1, obs2)

	assert.Exactly([]int{1, 2, 3}, values)
	assert.Exactly(errors.New("Bang"), err)
}

func TestPipeUnsubscribe(t *testing.T) {
	assert := assert.New(t)

	obs1 := New(0)
	obs2 := New(0)
	wait := make(chan bool, 1)

	done := Pipe(obs1, obs2)

	values := []int{}
	go func() {
		<-wait
		obs2.Unsubscribe()
		for v := range obs2.Out() {
			values = append(values, v.(int))
		}
	}()

	go func() {
		obs1.In() <- 0
		wait <- true
	Loop:
		for i := 1; i <= 3; i++ {
			select {
			case <-obs1.OnUnsubscribe():
				break Loop
			case obs1.In() <- i:
			}
		}
		obs1.Close()
	}()

	<-done

	assert.Exactly([]int{0}, values)
}
