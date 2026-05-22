package heartbeat

import "sync/atomic"

type Runtime struct {
	Beats atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Beat() {

	r.Beats.Add(1)
}

func (r *Runtime) Count() uint64 {

	return r.Beats.Load()
}
