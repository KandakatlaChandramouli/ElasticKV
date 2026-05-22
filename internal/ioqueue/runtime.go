package ioqueue

type Request struct {
	ID uint64
}

type Runtime struct {
	Queue []Request
}

func NewRuntime() *Runtime {

	return &Runtime{
		Queue: make([]Request, 0),
	}
}

func (r *Runtime) Submit(
	id uint64,
) {

	r.Queue = append(
		r.Queue,
		Request{
			ID: id,
		},
	)
}
