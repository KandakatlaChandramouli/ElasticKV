package schedulerqueue

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/schedulerqueue"
)

func BenchmarkSchedulerQueue(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Push(uint64(i))
	}

	if len(runtime.Tasks) == 0 {
		b.Fatal("queue failed")
	}
}
