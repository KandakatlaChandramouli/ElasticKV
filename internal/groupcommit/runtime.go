package groupcommit

import "sync/atomic"

type Runtime struct {
	Commits atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Commit(
	count uint64,
) {

	r.Commits.Add(
		count,
	)
}
