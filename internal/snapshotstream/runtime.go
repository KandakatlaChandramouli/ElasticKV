package snapshotstream

import "sync/atomic"

type Runtime struct {
	Bytes atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Send(
	size uint64,
) {

	r.Bytes.Add(size)
}
