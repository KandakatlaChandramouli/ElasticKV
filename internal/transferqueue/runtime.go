package transferqueue

type Transfer struct {
	ID uint64
}

type Runtime struct {
	Transfers []Transfer
}

func NewRuntime() *Runtime {

	return &Runtime{
		Transfers: make([]Transfer, 0),
	}
}

func (r *Runtime) Push(
	id uint64,
) {

	r.Transfers = append(
		r.Transfers,
		Transfer{
			ID: id,
		},
	)
}
