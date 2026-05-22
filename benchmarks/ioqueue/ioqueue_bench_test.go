package ioqueue

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/ioqueue"
)

func BenchmarkIOQueue(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Submit(
			uint64(i),
		)
	}

	if len(runtime.Queue) == 0 {
		b.Fatal("queue failed")
	}
}
