package vectorscan

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/vectorscan"
)

func BenchmarkVectorScan(
	b *testing.B,
) {

	data := make(
		[]uint64,
		100000,
	)

	for i := range data {
		data[i] = uint64(i)
	}

	target := uint64(77777)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		ok := engine.Scan(
			data,
			target,
		)

		if !ok {
			b.Fatal("scan failed")
		}
	}
}
