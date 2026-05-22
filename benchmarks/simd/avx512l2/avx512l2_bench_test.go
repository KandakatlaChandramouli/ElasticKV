package avx512l2

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/simd/avx512l2"
)

func BenchmarkAvx512L2(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ReportAllocs()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		if !runtime.Execute() {
			b.Fatal("execution failed")
		}
	}

	if runtime.Count() == 0 {
		b.Fatal("invalid runtime count")
	}
}
