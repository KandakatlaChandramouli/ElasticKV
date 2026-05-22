package fragmentation

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/fragmentation"
)

func BenchmarkFragmentation(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Allocate()
	}

	if runtime.Fragments == 0 {
		b.Fatal("fragmentation failed")
	}
}
