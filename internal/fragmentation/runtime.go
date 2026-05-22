package fragmentation

type Runtime struct {
	Fragments uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Allocate() {
	r.Fragments++
}
