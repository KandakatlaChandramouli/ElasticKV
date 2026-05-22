package lsm

import (
	"sort"
	"sync"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/sstable"
)

type MemTable struct {
	Data  map[uint64][]byte
	Keys  []uint64
	Size  uint64
	Mutex sync.RWMutex
}

func NewMemTable() *MemTable {

	return &MemTable{
		Data: make(map[uint64][]byte),
	}
}

func (m *MemTable) Put(
	key uint64,
	value []byte,
) {

	m.Mutex.Lock()

	defer m.Mutex.Unlock()

	_, exists := m.Data[key]

	if !exists {
		m.Keys = append(m.Keys, key)
	}

	m.Data[key] = value

	m.Size += uint64(len(value))
}

func (m *MemTable) Freeze() []sstable.Entry {

	m.Mutex.RLock()

	defer m.Mutex.RUnlock()

	sort.Slice(
		m.Keys,
		func(i, j int) bool {
			return m.Keys[i] < m.Keys[j]
		},
	)

	entries := make(
		[]sstable.Entry,
		0,
		len(m.Keys),
	)

	for _, key := range m.Keys {

		entries = append(
			entries,
			sstable.Entry{
				Key:   key,
				Value: m.Data[key],
			},
		)
	}

	return entries
}
