package sstable

import (
	"bytes"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/compaction"
	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/sstable"
)

func BenchmarkCompactionMerge(
	b *testing.B,
) {

	left := make([]engine.Entry, 0, 100000)
	right := make([]engine.Entry, 0, 100000)

	payload := bytes.Repeat([]byte("C"), 1024)

	for i := 0; i < 100000; i++ {

		left = append(
			left,
			engine.Entry{
				Key:   uint64(i * 2),
				Value: payload,
			},
		)

		right = append(
			right,
			engine.Entry{
				Key:   uint64(i*2 + 1),
				Value: payload,
			},
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		merged := compaction.Merge(
			left,
			right,
		)

		if len(merged) != 200000 {
			b.Fatal("merge failed")
		}
	}
}
