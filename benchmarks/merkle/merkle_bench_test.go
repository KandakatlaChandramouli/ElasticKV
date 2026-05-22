package merkle

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/merkle"
)

func BenchmarkMerkleHash(
	b *testing.B,
) {

	payload := bytes.Repeat(
		[]byte("M"),
		4096,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		hash := engine.Hash(
			payload,
		)

		if hash == 0 {
			b.Fatal("hash failed")
		}
	}
}
