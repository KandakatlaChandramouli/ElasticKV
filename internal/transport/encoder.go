package transport

import (
	"encoding/binary"
	"io"
)

func WriteFrame(
	writer io.Writer,
	frame ReplicationFrame,
) error {

	payloadLength := uint32(
		len(frame.Payload),
	)

	totalSize := uint32(
		8 + 4 + payloadLength,
	)

	header := make([]byte, 16)

	binary.LittleEndian.PutUint32(
		header[0:4],
		totalSize,
	)

	binary.LittleEndian.PutUint64(
		header[4:12],
		frame.SequenceID,
	)

	binary.LittleEndian.PutUint32(
		header[12:16],
		payloadLength,
	)

	_, err := writer.Write(header)

	if err != nil {
		return err
	}

	_, err = writer.Write(frame.Payload)

	return err
}
