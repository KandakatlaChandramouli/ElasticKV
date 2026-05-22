package failover

import "sync/atomic"

type Runtime struct {
	Switches atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Switch() {

	r.Switches.Add(1)
}
