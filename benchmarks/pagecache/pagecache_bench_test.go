package pagecache

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/pagecache"
)

func BenchmarkPageCacheLookup(
	b *testing.B,
) {

	cache := engine.New()

	payload := bytes.Repeat(
		[]byte("P"),
		4096,
	)

	for i := 0; i < 100000; i++ {

		cache.Put(
			uint64(i),
			payload,
		)
	}

	target := uint64(77777)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		value, ok := cache.Get(
			target,
		)

		if !ok {
			b.Fatal("lookup failed")
		}

		if len(value) != 4096 {
			b.Fatal("invalid page")
		}
	}
}
