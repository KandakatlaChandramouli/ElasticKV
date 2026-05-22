package checkpoint

import (
	"sync"
)

type Runtime struct {
	State    map[uint64][]byte
	Metadata Metadata
	Mutex    sync.RWMutex
}

func NewRuntime() *Runtime {

	return &Runtime{
		State: make(
			map[uint64][]byte,
		),
	}
}

func (r *Runtime) Apply(
	sequenceID uint64,
	payload []byte,
) {

	r.Mutex.Lock()

	defer r.Mutex.Unlock()

	r.State[sequenceID] = payload

	r.Metadata.LastSequenceID =
		sequenceID

	r.Metadata.EntryCount++
}

func (r *Runtime) Snapshot(
	path string,
) error {

	r.Mutex.RLock()

	defer r.Mutex.RUnlock()

	return SaveBinary(
		path,
		r.State,
		r.Metadata,
	)
}

func (r *Runtime) Restore(
	path string,
) error {

	r.Mutex.Lock()

	defer r.Mutex.Unlock()

	state,
		metadata,
		err := LoadBinary(path)

	if err != nil {
		return err
	}

	r.State = state

	r.Metadata = metadata

	return nil
}
