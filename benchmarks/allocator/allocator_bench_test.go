package allocator

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/allocator"
)

func BenchmarkAllocator(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		buffer := runtime.Allocate(
			4096,
		)

		if len(buffer) != 4096 {
			b.Fatal("allocation failed")
		}
	}
}
