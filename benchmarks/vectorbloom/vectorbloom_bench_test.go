package vectorbloom

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/vectorbloom"
)

func BenchmarkVectorBloom(
	b *testing.B,
) {

	data := make(
		[]uint64,
		100000,
	)

	for i := range data {
		data[i] = uint64(i)
	}

	target := uint64(88888)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		ok := engine.Contains(
			data,
			target,
		)

		if !ok {
			b.Fatal("lookup failed")
		}
	}
}
