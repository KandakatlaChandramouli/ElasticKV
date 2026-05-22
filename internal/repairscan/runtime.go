package repairscan

type Runtime struct {
	Scans uint64
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Scan() {
	r.Scans++
}
