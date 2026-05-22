package consistency

type Runtime struct {
	Reads  uint64
	Writes uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Read() {
	r.Reads++
}

func (r *Runtime) Write() {
	r.Writes++
}
