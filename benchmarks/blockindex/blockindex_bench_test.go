package blockindex

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/blockindex"
)

func BenchmarkBlockIndexLookup(
	b *testing.B,
) {

	index := engine.New()

	for i := 0; i < 100000; i++ {

		index.Add(
			uint64(i),
			uint64(i*4096),
		)
	}

	target := uint64(77777)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		offset, ok := index.Lookup(
			target,
		)

		if !ok {
			b.Fatal("lookup failed")
		}

		if offset == 0 {
			b.Fatal("invalid offset")
		}
	}
}
