package rdma

import (
	"bytes"
	"testing"
)

func BenchmarkZeroCopy(
	b *testing.B,
) {

	src := bytes.Repeat(
		[]byte("a"),
		1<<20,
	)

	dst := make([]byte, len(src))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		copy(dst, src)
	}
}
