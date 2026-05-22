package transport

import (
	"bytes"
	"encoding/binary"
	"io"
)

type ReplicationFrame struct {
	SequenceID uint64
	Payload    []byte
}

func EncodeFrame(
	frame ReplicationFrame,
) ([]byte, error) {

	payloadLength := uint32(
		len(frame.Payload),
	)

	totalSize := uint32(
		8 + 4 + payloadLength,
	)

	buf := new(bytes.Buffer)

	err := binary.Write(
		buf,
		binary.LittleEndian,
		totalSize,
	)

	if err != nil {
		return nil, err
	}

	err = binary.Write(
		buf,
		binary.LittleEndian,
		frame.SequenceID,
	)

	if err != nil {
		return nil, err
	}

	err = binary.Write(
		buf,
		binary.LittleEndian,
		payloadLength,
	)

	if err != nil {
		return nil, err
	}

	_, err = buf.Write(frame.Payload)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DecodeFrame(
	reader io.Reader,
) (ReplicationFrame, error) {

	var totalSize uint32

	err := binary.Read(
		reader,
		binary.LittleEndian,
		&totalSize,
	)

	if err != nil {
		return ReplicationFrame{}, err
	}

	var sequenceID uint64

	err = binary.Read(
		reader,
		binary.LittleEndian,
		&sequenceID,
	)

	if err != nil {
		return ReplicationFrame{}, err
	}

	var payloadLength uint32

	err = binary.Read(
		reader,
		binary.LittleEndian,
		&payloadLength,
	)

	if err != nil {
		return ReplicationFrame{}, err
	}

	payload := make([]byte, payloadLength)

	_, err = io.ReadFull(
		reader,
		payload,
	)

	if err != nil {
		return ReplicationFrame{}, err
	}

	return ReplicationFrame{
		SequenceID: sequenceID,
		Payload:    payload,
	}, nil
}
