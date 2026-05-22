package writebatch

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/writebatch"
)

func BenchmarkWriteBatch(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Put(
			uint64(i),
		)
	}

	if len(runtime.Entries) == 0 {
		b.Fatal("batch failed")
	}
}
