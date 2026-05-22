package crashsim

import "sync/atomic"

type Runtime struct {
	Crashes atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Crash() {

	r.Crashes.Add(1)
}
