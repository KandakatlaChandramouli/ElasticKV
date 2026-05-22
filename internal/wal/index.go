package wal

import (
	"sort"
	"sync"
)

type IndexEntry struct {
	SequenceID uint64
	Offset     uint64
	SegmentID  uint64
}

type SparseIndex struct {
	Entries  []IndexEntry
	Interval uint64
	Mutex    sync.RWMutex
}

func NewSparseIndex(
	interval uint64,
) *SparseIndex {

	return &SparseIndex{
		Entries:  make([]IndexEntry, 0),
		Interval: interval,
	}
}

func (s *SparseIndex) Add(
	sequenceID uint64,
	offset uint64,
	segmentID uint64,
) {

	if sequenceID%s.Interval != 0 {
		return
	}

	s.Mutex.Lock()

	defer s.Mutex.Unlock()

	s.Entries = append(
		s.Entries,
		IndexEntry{
			SequenceID: sequenceID,
			Offset:     offset,
			SegmentID:  segmentID,
		},
	)
}

func (s *SparseIndex) Find(
	sequenceID uint64,
) (IndexEntry, bool) {

	s.Mutex.RLock()

	defer s.Mutex.RUnlock()

	if len(s.Entries) == 0 {
		return IndexEntry{}, false
	}

	index := sort.Search(
		len(s.Entries),
		func(i int) bool {

			return s.Entries[i].SequenceID >= sequenceID
		},
	)

	if index >= len(s.Entries) {

		return s.Entries[len(s.Entries)-1], true
	}

	return s.Entries[index], true
}
