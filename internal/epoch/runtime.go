package epoch

import "sync/atomic"

type Runtime struct {
	Epoch atomic.Uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Advance() {
	r.Epoch.Add(1)
}
