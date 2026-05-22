package segmentgc

type Runtime struct {
	Count uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Execute() {
	r.Count++
}
