package snapshotter

import (
	"sync/atomic"
	"time"
)

type Runtime struct {
	Snapshots atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Snapshot() {

	time.Sleep(
		time.Microsecond,
	)

	r.Snapshots.Add(1)
}
