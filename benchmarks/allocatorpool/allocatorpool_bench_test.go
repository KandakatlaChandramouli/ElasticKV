package allocatorpool

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/allocatorpool"
)

func BenchmarkAllocatorPool(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		buffer := runtime.Get()

		runtime.Put(
			buffer,
		)
	}
}
