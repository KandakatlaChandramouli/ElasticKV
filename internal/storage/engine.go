package storage

import "os"

type Engine struct {
	file      *os.File
	data      []byte
	size      int
	allocator *Allocator
}

func Open(path string, size int) (*Engine, error) {

	f, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_RDWR,
		0644,
	)

	if err != nil {
		return nil, err
	}

	if err := f.Truncate(int64(size)); err != nil {
		return nil, err
	}

	data, err := mmapFile(f, size)

	if err != nil {
		return nil, err
	}

	return &Engine{
		file:      f,
		data:      data,
		size:      size,
		allocator: NewAllocator(),
	}, nil
}

func (e *Engine) Close() error {

	if err := munmap(e.data); err != nil {
		return err
	}

	return e.file.Close()
}

func (e *Engine) Write(
	sequenceID uint64,
	payload []byte,
) (uint64, error) {

	if len(payload) > MaxPayloadSize {
		return 0, ErrBlockTooLarge
	}

	offset := e.allocator.Allocate()

	if int(offset)+BlockSize > e.size {
		return 0, ErrOutOfBounds
	}

	block := e.data[offset : offset+BlockSize]

	header := BlockHeader{
		SequenceID: sequenceID,
		Length:     uint32(len(payload)),
		Flags:      0,
	}

	writeHeader(block[:HeaderSize], header)

	copy(block[HeaderSize:], payload)

	return offset, nil
}

func (e *Engine) Read(
	offset uint64,
) ([]byte, BlockHeader, error) {

	if int(offset)+BlockSize > e.size {
		return nil, BlockHeader{}, ErrOutOfBounds
	}

	block := e.data[offset : offset+BlockSize]

	header := readHeader(block[:HeaderSize])

	payload := make([]byte, header.Length)

	copy(
		payload,
		block[HeaderSize:HeaderSize+header.Length],
	)

	return payload, header, nil
}
