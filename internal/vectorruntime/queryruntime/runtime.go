package queryruntime

type Query struct {
	Vector []float32
	TopK   int
}

type Runtime struct {
	queries []Query
}

func NewRuntime() *Runtime {
	return &Runtime{
		queries: make([]Query, 0),
	}
}

func (r *Runtime) Submit(
	query Query,
) {

	r.queries = append(
		r.queries,
		query,
	)
}

func (r *Runtime) Count() int {
	return len(r.queries)
}
