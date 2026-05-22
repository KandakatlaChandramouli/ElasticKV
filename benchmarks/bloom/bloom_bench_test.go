package bloom

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/bloom"
)

func BenchmarkBloomLookup(
	b *testing.B,
) {

	filter := engine.New(
		1000000,
	)

	for i := 0; i < 100000; i++ {

		filter.Add(
			uint64(i),
		)
	}

	target := uint64(77777)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		ok := filter.MayContain(
			target,
		)

		if !ok {
			b.Fatal("lookup failed")
		}
	}
}
