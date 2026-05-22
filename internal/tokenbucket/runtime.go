package tokenbucket

import "sync/atomic"

type Runtime struct {
	Tokens atomic.Uint64
}

func NewRuntime(
	initial uint64,
) *Runtime {

	r := &Runtime{}
	r.Tokens.Store(initial)

	return r
}

func (r *Runtime) Consume() bool {

	for {
		current := r.Tokens.Load()

		if current == 0 {
			return false
		}

		if r.Tokens.CompareAndSwap(
			current,
			current-1,
		) {
			return true
		}
	}
}
