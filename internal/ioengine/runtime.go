package ioengine

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
