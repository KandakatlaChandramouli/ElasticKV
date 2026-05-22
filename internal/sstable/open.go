package sstable

import (
	"encoding/binary"
	"os"
	"syscall"
)

func Open(
	path string,
) (*Table, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	info, err := file.Stat()

	if err != nil {
		return nil, err
	}

	data, err := syscall.Mmap(
		int(file.Fd()),
		0,
		int(info.Size()),
		syscall.PROT_READ,
		syscall.MAP_SHARED,
	)

	if err != nil {
		return nil, err
	}

	count := binary.LittleEndian.Uint64(
		data[0:8],
	)

	index := make(
		[]IndexEntry,
		0,
		count,
	)

	offset := uint64(16)

	for i := uint64(0); i < count; i++ {

		key := binary.LittleEndian.Uint64(
			data[offset : offset+8],
		)

		length := binary.LittleEndian.Uint64(
			data[offset+8 : offset+16],
		)

		index = append(
			index,
			IndexEntry{
				Key:    key,
				Offset: offset + 16,
				Length: length,
			},
		)

		offset += 16 + length
	}

	return &Table{
		Path:  path,
		Index: index,
		Data:  data,
	}, nil
}
