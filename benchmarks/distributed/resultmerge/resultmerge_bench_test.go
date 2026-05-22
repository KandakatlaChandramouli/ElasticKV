package resultmerge

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/distributed/resultmerge"
)

func BenchmarkResultMerge(
	b *testing.B,
) {

	groups := [][]engine.Result{
		{
			{ID: 1, Score: 0.9},
			{ID: 2, Score: 0.8},
		},
		{
			{ID: 3, Score: 0.95},
			{ID: 4, Score: 0.7},
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = engine.Merge(
			groups,
			2,
		)
	}
}
