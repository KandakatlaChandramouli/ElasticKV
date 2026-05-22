package kernelqueue

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/gpuexecution/kernelqueue"
)

func BenchmarkKernelqueue(
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
		b.Fatal("invalid count")
	}
}
