package zerocopy

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/rdma/zerocopy"
)

func BenchmarkZeroCopy(
	b *testing.B,
) {

	src := make([]byte, 1<<20)
	dst := make([]byte, 1<<20)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = engine.Transfer(
			src,
			dst,
		)
	}
}
