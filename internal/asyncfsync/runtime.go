package asyncfsync

import "sync/atomic"

type Runtime struct {
	Syncs atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Sync() {

	go func() {
		r.Syncs.Add(1)
	}()
}
