package consistency

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/consistency"
)

func BenchmarkConsistency(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Read()
		runtime.Write()
	}

	if runtime.Reads == 0 {
		b.Fatal("consistency failed")
	}
}
