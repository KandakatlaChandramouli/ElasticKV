package lease

import "sync/atomic"

type Runtime struct {
	Leases atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Acquire() {

	r.Leases.Add(1)
}
