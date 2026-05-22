package snapshotqueue

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/snapshotqueue"
)

func BenchmarkSnapshotQueue(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Push(uint64(i))
	}

	if len(runtime.Queue) == 0 {
		b.Fatal("snapshot queue failed")
	}
}
