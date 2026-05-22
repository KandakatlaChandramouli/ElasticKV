package repair

import "sync/atomic"

type Runtime struct {
	Repairs atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Repair() {

	r.Repairs.Add(1)
}
