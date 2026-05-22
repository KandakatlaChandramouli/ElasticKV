package ratelimit

import "sync/atomic"

type Runtime struct {
	Count atomic.Uint64
	Limit uint64
}

func NewRuntime(
	limit uint64,
) *Runtime {

	return &Runtime{
		Limit: limit,
	}
}

func (r *Runtime) Allow() bool {

	current := r.Count.Add(1)

	return current <= r.Limit
}
