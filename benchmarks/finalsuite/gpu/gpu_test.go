package gpu

import (
	"testing"
)

func BenchmarkGPUKernel(
	b *testing.B,
) {

	buffer := make([]float32, 10_000_000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		for j := range buffer {
			buffer[j] += 1
		}
	}
}
