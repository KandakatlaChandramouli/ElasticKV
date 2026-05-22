package streaming

import "sync/atomic"

type Runtime struct {
	Bytes atomic.Uint64
}

func NewRuntime() *Runtime {

	return &Runtime{}
}

func (r *Runtime) Stream(
	size uint64,
) {

	r.Bytes.Add(size)
}
