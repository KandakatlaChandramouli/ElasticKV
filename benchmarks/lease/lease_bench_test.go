package lease

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/lease"
)

func BenchmarkLeaseAcquire(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Acquire()
	}

	if runtime.Leases.Load() == 0 {
		b.Fatal("lease failed")
	}
}
