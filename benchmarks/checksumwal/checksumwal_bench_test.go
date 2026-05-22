package checksumwal

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/checksumwal"
)

func BenchmarkChecksumWAL(
	b *testing.B,
) {

	payload := bytes.Repeat(
		[]byte("W"),
		4096,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		entry := engine.Encode(
			payload,
		)

		ok := engine.Verify(
			entry,
		)

		if !ok {
			b.Fatal("verification failed")
		}
	}
}
