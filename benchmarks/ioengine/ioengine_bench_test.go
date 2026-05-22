package ioengine

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/ioengine"
)

func BenchmarkIOEngine(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Execute()
	}

	if runtime.Operations.Load() == 0 {
		b.Fatal("io engine failed")
	}
}
