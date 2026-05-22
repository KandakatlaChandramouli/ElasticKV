package network

import "encoding/binary"

type Packet struct {
	Key   uint64
	Value []byte
}

func Encode(
	packet Packet,
) []byte {

	buffer := make(
		[]byte,
		16+len(packet.Value),
	)

	binary.LittleEndian.PutUint64(
		buffer[0:8],
		packet.Key,
	)

	binary.LittleEndian.PutUint64(
		buffer[8:16],
		uint64(len(packet.Value)),
	)

	copy(
		buffer[16:],
		packet.Value,
	)

	return buffer
}

func Decode(
	data []byte,
) Packet {

	key := binary.LittleEndian.Uint64(
		data[0:8],
	)

	length := binary.LittleEndian.Uint64(
		data[8:16],
	)

	value := data[16 : 16+length]

	return Packet{
		Key:   key,
		Value: value,
	}
}
