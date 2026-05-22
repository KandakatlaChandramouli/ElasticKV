package txn

import "sync"

type Version struct {
	Timestamp uint64
	Value     []byte
}

type Store struct {
	Data  map[uint64][]Version
	Mutex sync.RWMutex
}

func NewStore() *Store {

	return &Store{
		Data: make(map[uint64][]Version),
	}
}

func (s *Store) Put(
	key uint64,
	ts uint64,
	value []byte,
) {

	s.Mutex.Lock()

	defer s.Mutex.Unlock()

	s.Data[key] = append(
		s.Data[key],
		Version{
			Timestamp: ts,
			Value:     value,
		},
	)
}

func (s *Store) Get(
	key uint64,
	ts uint64,
) ([]byte, bool) {

	s.Mutex.RLock()

	defer s.Mutex.RUnlock()

	versions, ok := s.Data[key]

	if !ok {
		return nil, false
	}

	for i := len(versions) - 1; i >= 0; i-- {

		version := versions[i]

		if version.Timestamp <= ts {
			return version.Value, true
		}
	}

	return nil, false
}
