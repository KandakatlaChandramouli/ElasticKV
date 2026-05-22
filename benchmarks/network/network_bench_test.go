package network

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/network"
)

func BenchmarkPacketEncode(
	b *testing.B,
) {

	payload := bytes.Repeat(
		[]byte("N"),
		1024,
	)

	packet := engine.Packet{
		Key:   777,
		Value: payload,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		data := engine.Encode(
			packet,
		)

		if len(data) == 0 {
			b.Fatal("encode failed")
		}
	}
}

func BenchmarkPacketDecode(
	b *testing.B,
) {

	payload := bytes.Repeat(
		[]byte("D"),
		1024,
	)

	packet := engine.Packet{
		Key:   999,
		Value: payload,
	}

	data := engine.Encode(packet)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		result := engine.Decode(
			data,
		)

		if result.Key != 999 {
			b.Fatal("decode failed")
		}
	}
}

func BenchmarkBufferPool(
	b *testing.B,
) {

	pool := engine.NewPool(
		2048,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		buffer := pool.Get()

		if len(buffer) != 2048 {
			b.Fatal("invalid buffer")
		}

		pool.Put(buffer)
	}
}
