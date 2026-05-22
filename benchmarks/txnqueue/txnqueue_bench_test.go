package txnqueue

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/txnqueue"
)

func BenchmarkTxnQueue(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Enqueue(uint64(i))
	}

	if len(runtime.Queue) == 0 {
		b.Fatal("txn queue failed")
	}
}
