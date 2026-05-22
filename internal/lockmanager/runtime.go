package lockmanager

type Runtime struct {
	Locks map[uint64]bool
}

func NewRuntime() *Runtime {

	return &Runtime{
		Locks: make(map[uint64]bool),
	}
}

func (r *Runtime) Lock(
	id uint64,
) {

	r.Locks[id] = true
}
