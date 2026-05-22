package checkpoint

import (
	"github.com/KandakatlaChandramouli/ElasticKV/internal/filter"
)

type IndexedSnapshot struct {
	Snapshot *MMapSnapshot
	Index    *SnapshotIndex
	Filter   *filter.BloomFilter
}

func OpenIndexedSnapshot(
	path string,
) (
	*IndexedSnapshot,
	error,
) {

	snapshot, err := LoadMMap(path)

	if err != nil {
		return nil, err
	}

	index := BuildIndex(snapshot.Data)

	bloom := filter.NewBloomFilter(1000000)

	for _, entry := range index.Entries {
		bloom.Add(entry.Key)
	}

	return &IndexedSnapshot{
		Snapshot: snapshot,
		Index:    index,
		Filter:   bloom,
	}, nil
}

func (i *IndexedSnapshot) Lookup(
	key uint64,
) ([]byte, bool) {

	if !i.Filter.MayContain(key) {
		return nil, false
	}

	entry, ok := i.Index.Lookup(key)

	if !ok {
		return nil, false
	}

	value := i.Snapshot.Data[entry.Offset : entry.Offset+entry.Length]

	return value, true
}

func (i *IndexedSnapshot) Close() error {
	return i.Snapshot.Close()
}
