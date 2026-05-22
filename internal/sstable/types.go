package sstable

type Entry struct {
	Key   uint64
	Value []byte
}

type IndexEntry struct {
	Key    uint64
	Offset uint64
	Length uint64
}
