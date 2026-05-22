package streamqueue

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/streamqueue"
)

func BenchmarkStreamQueue(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Push(uint64(i))
	}

	if len(runtime.Entries) == 0 {
		b.Fatal("stream queue failed")
	}
}
