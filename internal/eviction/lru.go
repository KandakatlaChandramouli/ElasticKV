package eviction

import "container/list"

type Entry struct {
	Key uint64
}

type Runtime struct {
	List *list.List
}

func NewRuntime() *Runtime {

	return &Runtime{
		List: list.New(),
	}
}

func (r *Runtime) Touch(
	key uint64,
) {

	r.List.PushFront(
		Entry{
			Key: key,
		},
	)
}

func (r *Runtime) Evict() bool {

	if r.List.Len() == 0 {

		r.Touch(0)
	}

	back := r.List.Back()

	if back == nil {
		return false
	}

	r.List.Remove(
		back,
	)

	return true
}
