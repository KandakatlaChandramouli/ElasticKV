package recall

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/hnsw/core"
)

func BenchmarkRecall(
	b *testing.B,
) {

	index := engine.New()

	vector := make([]float32, 768)

	for i := 0; i < 50000; i++ {

		index.Insert(
			i,
			vector,
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		results := index.Search(
			vector,
			100,
		)

		if len(results) == 0 {
			b.Fatal("no results")
		}
	}
}
