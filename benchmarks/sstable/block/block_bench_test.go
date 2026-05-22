package block

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/sstable/block"
)

func BenchmarkBlock(
	b *testing.B,
) {

	block := engine.New()

	value := make([]byte, 256)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		block.Add(
			"key",
			value,
		)
	}
}
