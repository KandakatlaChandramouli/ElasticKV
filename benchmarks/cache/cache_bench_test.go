package cache

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/cache"
)

func BenchmarkShardedCache(
	b *testing.B,
) {

	cache := engine.NewSharded(
		64,
		100000,
	)

	payload := bytes.Repeat(
		[]byte("C"),
		1024,
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
			b.Fatal("cache miss")
		}

		if len(value) != 1024 {
			b.Fatal("invalid payload")
		}
	}
}
