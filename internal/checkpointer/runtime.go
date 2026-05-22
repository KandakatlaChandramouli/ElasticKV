package checkpointer

import "sync/atomic"

type Runtime struct {
	Points atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Checkpoint() {

	r.Points.Add(1)
}
