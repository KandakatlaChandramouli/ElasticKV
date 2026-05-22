package allocator

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/gpukernel/allocator"
)

func BenchmarkAllocator(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		if !runtime.Execute() {
			b.Fatal("execution failed")
		}
	}
}
