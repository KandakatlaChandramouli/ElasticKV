package segmentindex

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/segmentindex"
)

func BenchmarkSegmentIndex(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Put(uint64(i))
	}

	if len(runtime.Segments) == 0 {
		b.Fatal("segment index failed")
	}
}
