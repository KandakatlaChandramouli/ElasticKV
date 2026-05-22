package wal

import (
	"encoding/binary"
	"os"
	"sync"
)

type Segment struct {
	File   *os.File
	Mutex  sync.Mutex
	Offset uint64
}

func OpenSegment(
	path string,
) (*Segment, error) {

	file, err := os.OpenFile(
		path,
		os.O_CREATE|
			os.O_RDWR|
			os.O_APPEND,
		0644,
	)

	if err != nil {
		return nil, err
	}

	info, err := file.Stat()

	if err != nil {
		return nil, err
	}

	return &Segment{
		File:   file,
		Offset: uint64(info.Size()),
	}, nil
}

func (s *Segment) Append(
	entry Entry,
) (uint64, error) {

	s.Mutex.Lock()

	defer s.Mutex.Unlock()

	offset := s.Offset

	payloadLength := uint64(
		len(entry.Payload),
	)

	header := make([]byte, HeaderSize)

	binary.LittleEndian.PutUint64(
		header[0:8],
		entry.SequenceID,
	)

	binary.LittleEndian.PutUint64(
		header[8:16],
		payloadLength,
	)

	_, err := s.File.Write(header)

	if err != nil {
		return 0, err
	}

	_, err = s.File.Write(
		entry.Payload,
	)

	if err != nil {
		return 0, err
	}

	s.Offset += HeaderSize + payloadLength

	return offset, nil
}

func (s *Segment) Sync() error {
	return s.File.Sync()
}

func (s *Segment) Close() error {
	return s.File.Close()
}
