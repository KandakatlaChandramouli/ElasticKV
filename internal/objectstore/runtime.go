package objectstore

type Object struct {
	ID uint64
}

type Runtime struct {
	Objects map[uint64]Object
}

func NewRuntime() *Runtime {

	return &Runtime{
		Objects: make(map[uint64]Object),
	}
}

func (r *Runtime) Put(
	id uint64,
) {

	r.Objects[id] = Object{
		ID: id,
	}
}
