package txnqueue

type Transaction struct {
	ID uint64
}

type Runtime struct {
	Queue []Transaction
}

func NewRuntime() *Runtime {

	return &Runtime{
		Queue: make([]Transaction, 0),
	}
}

func (r *Runtime) Enqueue(
	id uint64,
) {

	r.Queue = append(
		r.Queue,
		Transaction{
			ID: id,
		},
	)
}
