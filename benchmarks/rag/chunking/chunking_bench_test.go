package chunking

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/rag/chunking"
)

func BenchmarkChunking(
	b *testing.B,
) {

	tokens := make([]string, 10000)

	for i := range tokens {
		tokens[i] = "token"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = engine.Chunk(
			tokens,
			256,
		)
	}
}
