package vectorcache

type Vector struct {
	ID uint64
}

type Runtime struct {
	Cache map[uint64]Vector
}

func NewRuntime() *Runtime {

	return &Runtime{
		Cache: make(map[uint64]Vector),
	}
}

func (r *Runtime) Put(
	id uint64,
) {

	r.Cache[id] = Vector{
		ID: id,
	}
}
