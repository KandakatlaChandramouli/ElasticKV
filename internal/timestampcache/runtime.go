package timestampcache

import "sync/atomic"

type Runtime struct {
	Timestamp atomic.Uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Advance() {
	r.Timestamp.Add(1)
}
