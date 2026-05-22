package fencepointer

type Pointer struct {
	Key    uint64
	Offset uint64
}

type Runtime struct {
	Pointers []Pointer
}

func NewRuntime() *Runtime {

	return &Runtime{
		Pointers: make([]Pointer, 0),
	}
}

func (r *Runtime) Add(
	key uint64,
	offset uint64,
) {

	r.Pointers = append(
		r.Pointers,
		Pointer{
			Key:    key,
			Offset: offset,
		},
	)
}

func (r *Runtime) Search(
	key uint64,
) (uint64, bool) {

	for _, pointer := range r.Pointers {

		if pointer.Key == key {
			return pointer.Offset, true
		}
	}

	return 0, false
}
