package checkpoint

import (
	"encoding/binary"
	"os"
	"syscall"
)

type MMapSnapshot struct {
	Data     []byte
	Metadata Metadata
}

func LoadMMap(
	path string,
) (*MMapSnapshot, error) {

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

	metadata := Metadata{
		LastSequenceID: binary.LittleEndian.Uint64(data[0:8]),
		WALSegmentID:   binary.LittleEndian.Uint64(data[8:16]),
		EntryCount:     binary.LittleEndian.Uint64(data[16:24]),
	}

	return &MMapSnapshot{
		Data:     data,
		Metadata: metadata,
	}, nil
}

func (m *MMapSnapshot) Iterate(
	handler func(
		key uint64,
		value []byte,
	) error,
) error {

	offset := HeaderSize

	for {

		if offset+16 >= len(m.Data) {
			break
		}

		key := binary.LittleEndian.Uint64(
			m.Data[offset : offset+8],
		)

		length := binary.LittleEndian.Uint64(
			m.Data[offset+8 : offset+16],
		)

		offset += 16

		if offset+int(length) > len(m.Data) {
			break
		}

		value := m.Data[offset : offset+int(length)]

		err := handler(
			key,
			value,
		)

		if err != nil {
			return err
		}

		offset += int(length)
	}

	return nil
}

func (m *MMapSnapshot) Close() error {

	return syscall.Munmap(
		m.Data,
	)
}
