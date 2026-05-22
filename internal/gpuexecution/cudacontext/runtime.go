package cudacontext

import "sync/atomic"

type Runtime struct {
	Operations atomic.Uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Execute() bool {
	r.Operations.Add(1)
	return true
}

func (r *Runtime) Count() uint64 {
	return r.Operations.Load()
}
