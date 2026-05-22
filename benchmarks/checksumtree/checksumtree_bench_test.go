package checksumtree

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/checksumtree"
)

func BenchmarkChecksumTree(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	payload := bytes.Repeat(
		[]byte("C"),
		4096,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		sum := runtime.Sum(
			payload,
		)

		if sum == 0 {
			b.Fatal("checksum failed")
		}
	}
}
