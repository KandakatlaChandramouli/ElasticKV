package transport

import (
	"bytes"
	"encoding/binary"
)

type ReplicationFrame struct {
	SequenceID uint64
	Payload    []byte
}

func EncodeFrame(
	frame ReplicationFrame,
) ([]byte, error) {

	buf := new(bytes.Buffer)

	if err := binary.Write(
		buf,
		binary.LittleEndian,
		frame.SequenceID,
	); err != nil {
		return nil, err
	}

	payloadLength := uint32(
		len(frame.Payload),
	)

	if err := binary.Write(
		buf,
		binary.LittleEndian,
		payloadLength,
	); err != nil {
		return nil, err
	}

	_, err := buf.Write(frame.Payload)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DecodeFrame(
	data []byte,
) (ReplicationFrame, error) {

	reader := bytes.NewReader(data)

	var sequenceID uint64

	err := binary.Read(
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

	_, err = reader.Read(payload)

	if err != nil {
		return ReplicationFrame{}, err
	}

	return ReplicationFrame{
		SequenceID: sequenceID,
		Payload:    payload,
	}, nil
}
