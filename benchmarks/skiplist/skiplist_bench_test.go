package skiplist

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/skiplist"
)

func BenchmarkSkiplistLookup(
	b *testing.B,
) {

	list := engine.New()

	payload := bytes.Repeat(
		[]byte("S"),
		1024,
	)

	for i := 0; i < 100000; i++ {

		list.Insert(
			uint64(i),
			payload,
		)
	}

	target := uint64(77777)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		value, ok := list.Lookup(
			target,
		)

		if !ok {
			b.Fatal("lookup failed")
		}

		if len(value) != 1024 {
			b.Fatal("invalid payload")
		}
	}
}
