package telemetry

import "sync/atomic"

type Runtime struct {
	Events atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Record() {

	r.Events.Add(1)
}
