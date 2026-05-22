package compactor

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/compactor"
	"github.com/KandakatlaChandramouli/ElasticKV/internal/sstable"
)

func BenchmarkCompactorRuntime(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	left := make(
		[]sstable.Entry,
		0,
		100000,
	)

	right := make(
		[]sstable.Entry,
		0,
		100000,
	)

	payload := bytes.Repeat(
		[]byte("C"),
		1024,
	)

	for i := 0; i < 100000; i++ {

		left = append(
			left,
			sstable.Entry{
				Key:   uint64(i * 2),
				Value: payload,
			},
		)

		right = append(
			right,
			sstable.Entry{
				Key:   uint64(i*2 + 1),
				Value: payload,
			},
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		result := runtime.Compact(
			left,
			right,
		)

		if len(result) != 200000 {
			b.Fatal("compaction failed")
		}
	}
}
