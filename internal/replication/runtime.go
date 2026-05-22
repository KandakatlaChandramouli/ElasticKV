package replication

import "sync/atomic"

type Runtime struct {
	Replicas atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Replicate(
	count uint64,
) {

	r.Replicas.Add(count)
}
