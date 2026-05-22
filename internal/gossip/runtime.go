package gossip

import "sync/atomic"

type Runtime struct {
	Messages atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Broadcast() {

	r.Messages.Add(1)
}
