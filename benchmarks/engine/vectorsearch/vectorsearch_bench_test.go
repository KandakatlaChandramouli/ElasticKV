package vectorsearch

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/engine/vectorsearch"
)

func BenchmarkVectorSearch(
	b *testing.B,
) {

	scores := make([]float32, 10000)

	for i := range scores {
		scores[i] = float32(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = engine.Search(
			scores,
			10,
		)
	}
}
