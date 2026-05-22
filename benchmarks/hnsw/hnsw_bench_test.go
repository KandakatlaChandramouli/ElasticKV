package hnsw

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/hnsw/core"
)

func BenchmarkHNSWSearch(
	b *testing.B,
) {

	index := engine.New()

	vector := make([]float32, 768)

	for i := 0; i < 10000; i++ {

		index.Insert(
			i,
			vector,
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = index.Search(
			vector,
			10,
		)
	}
}
