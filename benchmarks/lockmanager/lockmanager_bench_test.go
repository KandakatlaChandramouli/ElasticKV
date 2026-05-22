package lockmanager

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/lockmanager"
)

func BenchmarkLockManager(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Lock(uint64(i))
	}

	if len(runtime.Locks) == 0 {
		b.Fatal("lock manager failed")
	}
}
