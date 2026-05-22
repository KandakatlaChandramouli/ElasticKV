package tombstone

import "sync"

type Runtime struct {
	Deleted map[uint64]struct{}
	Mutex   sync.RWMutex
}

func NewRuntime() *Runtime {

	return &Runtime{
		Deleted: make(
			map[uint64]struct{},
		),
	}
}

func (r *Runtime) Delete(
	key uint64,
) {

	r.Mutex.Lock()

	defer r.Mutex.Unlock()

	r.Deleted[key] = struct{}{}
}

func (r *Runtime) IsDeleted(
	key uint64,
) bool {

	r.Mutex.RLock()

	defer r.Mutex.RUnlock()

	_, ok := r.Deleted[key]

	return ok
}
