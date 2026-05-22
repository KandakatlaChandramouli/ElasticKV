package wal

import (
	"encoding/binary"
	"io"
	"os"
)

func ReplayFromOffset(
	path string,
	startOffset uint64,
	handler func(Entry) error,
) error {

	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Seek(
		int64(startOffset),
		io.SeekStart,
	)

	if err != nil {
		return err
	}

	for {

		header := make([]byte, HeaderSize)

		_, err := io.ReadFull(
			file,
			header,
		)

		if err != nil {

			if err == io.EOF ||
				err == io.ErrUnexpectedEOF {

				return nil
			}

			return err
		}

		sequenceID := binary.LittleEndian.Uint64(
			header[0:8],
		)

		payloadLength := binary.LittleEndian.Uint64(
			header[8:16],
		)

		payload := make(
			[]byte,
			payloadLength,
		)

		_, err = io.ReadFull(
			file,
			payload,
		)

		if err != nil {
			return err
		}

		entry := Entry{
			SequenceID: sequenceID,
			Payload:    payload,
		}

		err = handler(entry)

		if err != nil {
			return err
		}
	}
}
