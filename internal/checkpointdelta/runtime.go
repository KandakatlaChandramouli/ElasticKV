package checkpointdelta

import "sync/atomic"

type Runtime struct {
	Deltas atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) ApplyDelta(
	count uint64,
) {

	r.Deltas.Add(count)
}
