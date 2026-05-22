package checkpoint

import (
	"encoding/binary"
	"sort"
)

type SnapshotIndexEntry struct {
	Key    uint64
	Offset uint64
	Length uint64
}

type SnapshotIndex struct {
	Entries []SnapshotIndexEntry
}

func BuildIndex(
	data []byte,
) *SnapshotIndex {

	index := &SnapshotIndex{
		Entries: make(
			[]SnapshotIndexEntry,
			0,
			1024,
		),
	}

	offset := uint64(HeaderSize)

	for {

		if offset+16 >= uint64(len(data)) {
			break
		}

		key := binary.LittleEndian.Uint64(
			data[offset : offset+8],
		)

		length := binary.LittleEndian.Uint64(
			data[offset+8 : offset+16],
		)

		valueOffset := offset + 16

		index.Entries = append(
			index.Entries,
			SnapshotIndexEntry{
				Key:    key,
				Offset: valueOffset,
				Length: length,
			},
		)

		offset =
			valueOffset + length
	}

	sort.Slice(
		index.Entries,
		func(i, j int) bool {

			return index.Entries[i].Key <
				index.Entries[j].Key
		},
	)

	return index
}

func (s *SnapshotIndex) Lookup(
	key uint64,
) (
	SnapshotIndexEntry,
	bool,
) {

	low := 0

	high := len(s.Entries) - 1

	for low <= high {

		mid := (low + high) / 2

		entry := s.Entries[mid]

		if entry.Key == key {
			return entry, true
		}

		if entry.Key < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return SnapshotIndexEntry{},
		false
}
