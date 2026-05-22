package clocksource

import "sync/atomic"

type Runtime struct {
	Ticks atomic.Uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Tick() {
	r.Ticks.Add(1)
}
