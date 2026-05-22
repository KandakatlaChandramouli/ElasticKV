package histogram

type Histogram struct {
	values []int64
}

func New() *Histogram {

	return &Histogram{
		values: make([]int64, 0),
	}
}

func (h *Histogram) Observe(
	value int64,
) {

	h.values = append(
		h.values,
		value,
	)
}

func (h *Histogram) Count() int {
	return len(h.values)
}
