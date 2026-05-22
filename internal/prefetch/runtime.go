package prefetch

import "sync/atomic"

type Runtime struct {
	Prefetches atomic.Uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Fetch() {
	r.Prefetches.Add(1)
}
