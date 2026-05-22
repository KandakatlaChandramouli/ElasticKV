package compression

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/compression"
)

func BenchmarkCompression(
	b *testing.B,
) {

	payload := bytes.Repeat(
		[]byte("Z"),
		1024*64,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		compressed, err := engine.Compress(
			payload,
		)

		if err != nil {
			b.Fatal(err)
		}

		decompressed, err := engine.Decompress(
			compressed,
		)

		if err != nil {
			b.Fatal(err)
		}

		if len(decompressed) != len(payload) {
			b.Fatal("decompression failed")
		}
	}
}
