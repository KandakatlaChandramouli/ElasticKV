package leader

import "sync/atomic"

type Runtime struct {
	Leader atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Elect(
	node uint64,
) {

	r.Leader.Store(
		node,
	)
}

func (r *Runtime) Current() uint64 {

	return r.Leader.Load()
}
