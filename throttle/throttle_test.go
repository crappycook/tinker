package throttle

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThrottle(t *testing.T) {
	var counter uint32
	r := NewValve(3)
	for i := 0; i < 10000; i++ {
		r.Add()
		go func(c *uint32) {
			defer r.Done()
			atomic.AddUint32(c, 1)
		}(&counter)
	}
	r.Wait()
	assert.Equal(t, 10000, int(counter))
}

func TestThrottleRun(t *testing.T) {
	var counter uint32
	r := NewValve(10)
	for i := 0; i < 10000; i++ {
		r.Run(func() {
			atomic.AddUint32(&counter, 1)
		})
	}
	r.Wait()
	assert.Equal(t, 10000, int(counter))
}
