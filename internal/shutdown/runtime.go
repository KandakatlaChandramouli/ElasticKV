package shutdown

import "sync/atomic"

type Runtime struct {
	Closed atomic.Bool
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Stop() {
	r.Closed.Store(true)
}

func (r *Runtime) IsStopped() bool {
	return r.Closed.Load()
}
