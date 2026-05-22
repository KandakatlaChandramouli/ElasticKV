package segment

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/segment"
)

func BenchmarkSegmentCreation(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Create(
			uint64(i),
			1024,
		)
	}

	if len(runtime.Segments) == 0 {
		b.Fatal("segment failed")
	}
}
