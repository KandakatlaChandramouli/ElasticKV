package checkpoint

import (
	"encoding/binary"
	"io"
	"os"
)

const (
	HeaderSize = 24
)

func SaveBinary(
	path string,
	state map[uint64][]byte,
	metadata Metadata,
) error {

	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	header := make([]byte, HeaderSize)

	binary.LittleEndian.PutUint64(
		header[0:8],
		metadata.LastSequenceID,
	)

	binary.LittleEndian.PutUint64(
		header[8:16],
		metadata.WALSegmentID,
	)

	binary.LittleEndian.PutUint64(
		header[16:24],
		metadata.EntryCount,
	)

	_, err = file.Write(header)

	if err != nil {
		return err
	}

	for key, value := range state {

		entryHeader := make([]byte, 16)

		binary.LittleEndian.PutUint64(
			entryHeader[0:8],
			key,
		)

		binary.LittleEndian.PutUint64(
			entryHeader[8:16],
			uint64(len(value)),
		)

		_, err = file.Write(entryHeader)

		if err != nil {
			return err
		}

		_, err = file.Write(value)

		if err != nil {
			return err
		}
	}

	return file.Sync()
}

func LoadBinary(
	path string,
) (
	map[uint64][]byte,
	Metadata,
	error,
) {

	file, err := os.Open(path)

	if err != nil {

		return nil,
			Metadata{},
			err
	}

	defer file.Close()

	header := make([]byte, HeaderSize)

	_, err = io.ReadFull(
		file,
		header,
	)

	if err != nil {

		return nil,
			Metadata{},
			err
	}

	metadata := Metadata{
		LastSequenceID: binary.LittleEndian.Uint64(
			header[0:8],
		),
		WALSegmentID: binary.LittleEndian.Uint64(
			header[8:16],
		),
		EntryCount: binary.LittleEndian.Uint64(
			header[16:24],
		),
	}

	state := make(
		map[uint64][]byte,
	)

	for {

		entryHeader := make([]byte, 16)

		_, err := io.ReadFull(
			file,
			entryHeader,
		)

		if err != nil {

			if err == io.EOF ||
				err == io.ErrUnexpectedEOF {

				break
			}

			return nil,
				Metadata{},
				err
		}

		key := binary.LittleEndian.Uint64(
			entryHeader[0:8],
		)

		length := binary.LittleEndian.Uint64(
			entryHeader[8:16],
		)

		value := make(
			[]byte,
			length,
		)

		_, err = io.ReadFull(
			file,
			value,
		)

		if err != nil {

			return nil,
				Metadata{},
				err
		}

		state[key] = value
	}

	return state,
		metadata,
		nil
}
