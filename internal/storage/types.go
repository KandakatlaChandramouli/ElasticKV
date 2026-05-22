package storage

const (
	BlockSize  = 4096
	HeaderSize = 16
)

type BlockHeader struct {
	SequenceID uint64
	Length     uint32
	Flags      uint32
}
