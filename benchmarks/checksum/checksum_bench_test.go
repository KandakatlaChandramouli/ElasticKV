package checksum

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/checksum"
)

func BenchmarkCRC32(
	b *testing.B,
) {

	payload := bytes.Repeat(
		[]byte("C"),
		4096,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		sum := engine.Compute(
			payload,
		)

		ok := engine.Verify(
			payload,
			sum,
		)

		if !ok {
			b.Fatal("checksum failed")
		}
	}
}
