package streamscheduler

import "sync/atomic"

type Runtime struct {
	Ops atomic.Uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Execute() {
	r.Ops.Add(1)
}

func (r *Runtime) Count() uint64 {
	return r.Ops.Load()
}
