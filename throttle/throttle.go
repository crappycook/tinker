package throttle

import (
	"sync"
)

type Valve struct {
	capacity uint64
	c        chan struct{}
	wg       *sync.WaitGroup
}

// 最多允许并行调度的数目
// 控制并发数量 防止并发数过高 把下游服务搞挂
func NewThrottle(capacity uint64) *Valve {
	return &Valve{
		capacity: capacity,
		c:        make(chan struct{}, capacity),
		wg:       new(sync.WaitGroup),
	}
}

func (v *Valve) Add() {
	select {
	case v.c <- struct{}{}:
		break
	}
	v.wg.Add(1)
}

func (v *Valve) Done() {
	<-v.c
	v.wg.Done()
}

func (v *Valve) Wait() {
	v.wg.Wait()
}

func (v *Valve) Run(fn func()) {
	v.Add()
	go func() {
		fn()
		v.Done()
	}()
}
