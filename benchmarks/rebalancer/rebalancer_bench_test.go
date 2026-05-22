package rebalancer

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/rebalancer"
)

func BenchmarkRebalancer(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Move(
			uint64(i),
			uint64(i%8),
		)
	}

	if len(runtime.Shards) == 0 {
		b.Fatal("rebalance failed")
	}
}
