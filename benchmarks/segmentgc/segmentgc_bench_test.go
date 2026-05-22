package segmentgc

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/segmentgc"
)

func BenchmarkSegmentGC(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Execute()
	}

	if runtime.Count == 0 {
		b.Fatal("runtime failed")
	}
}
