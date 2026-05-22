package vectorized

import "sync/atomic"

type Runtime struct {
	Ops atomic.Uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Execute() bool {
	r.Ops.Add(1)
	return true
}

func (r *Runtime) Count() uint64 {
	return r.Ops.Load()
}
