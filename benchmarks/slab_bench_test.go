package benchmarks

import (
	"bytes"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/storage"
)

func BenchmarkWrite(b *testing.B) {

	engine, err := storage.Open(
		"bench.db",
		1024*1024*1024,
	)

	if err != nil {
		b.Fatal(err)
	}

	defer engine.Close()

	payload := bytes.Repeat([]byte("A"), 1024)

	maxWrites := (1024 * 1024 * 1024) / storage.BlockSize

	b.ResetTimer()

	for i := 0; i < b.N && i < maxWrites; i++ {

		_, err := engine.Write(
			uint64(i),
			payload,
		)

		if err != nil {
			b.Fatal(err)
		}
	}
}
