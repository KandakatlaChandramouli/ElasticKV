package transformindex

import "sync/atomic"

type Runtime struct {
	Operations atomic.Uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Execute() {
	r.Operations.Add(1)
}

func (r *Runtime) Count() uint64 {
	return r.Operations.Load()
}
