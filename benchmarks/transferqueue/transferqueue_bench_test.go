package transferqueue

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/transferqueue"
)

func BenchmarkTransferQueue(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Push(uint64(i))
	}

	if len(runtime.Transfers) == 0 {
		b.Fatal("transfer queue failed")
	}
}
