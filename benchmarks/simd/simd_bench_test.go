package simd

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/simd"
)

func BenchmarkSIMDSum(
	b *testing.B,
) {

	data := make(
		[]uint64,
		100000,
	)

	for i := range data {
		data[i] = uint64(i)
	}

	var total uint64

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		total = engine.Sum(
			data,
		)
	}

	if total == 0 {
		b.Fatal("sum failed")
	}
}
