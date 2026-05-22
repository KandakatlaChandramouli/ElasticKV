package batchscheduler

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/batchscheduler"
)

func BenchmarkBatchScheduler(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Schedule(uint64(i))
	}

	if len(runtime.Batches) == 0 {
		b.Fatal("batch scheduler failed")
	}
}
