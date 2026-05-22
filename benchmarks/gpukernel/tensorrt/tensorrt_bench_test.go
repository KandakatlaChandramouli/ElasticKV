package tensorrt

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/gpukernel/tensorrt"
)

func BenchmarkTensorrt(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		if !runtime.Execute() {
			b.Fatal("execution failed")
		}
	}
}
