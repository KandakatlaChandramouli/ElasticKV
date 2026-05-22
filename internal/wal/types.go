package wal

const (
	HeaderSize = 16
)

type Entry struct {
	SequenceID uint64
	Payload    []byte
}
