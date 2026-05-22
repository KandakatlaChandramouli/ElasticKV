package storage

import "encoding/binary"

func writeHeader(dst []byte, h BlockHeader) {
	binary.LittleEndian.PutUint64(dst[0:8], h.SequenceID)
	binary.LittleEndian.PutUint32(dst[8:12], h.Length)
	binary.LittleEndian.PutUint32(dst[12:16], h.Flags)
}

func readHeader(src []byte) BlockHeader {
	return BlockHeader{
		SequenceID: binary.LittleEndian.Uint64(src[0:8]),
		Length:     binary.LittleEndian.Uint32(src[8:12]),
		Flags:      binary.LittleEndian.Uint32(src[12:16]),
	}
}
