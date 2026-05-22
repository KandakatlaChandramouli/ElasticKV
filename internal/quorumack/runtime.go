package quorumack

import "sync/atomic"

type Runtime struct {
	Required uint64
	Acks     atomic.Uint64
}

func NewRuntime(
	required int,
) *Runtime {

	return &Runtime{
		Required: uint64(required),
	}
}

func (r *Runtime) Ack() {

	r.Acks.Add(1)
}

func (r *Runtime) Quorum() bool {

	return r.Acks.Load() >=
		r.Required
}
