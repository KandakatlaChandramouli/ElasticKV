package checkpoint

type Metadata struct {
	LastSequenceID uint64
	WALSegmentID   uint64
	EntryCount     uint64
}
