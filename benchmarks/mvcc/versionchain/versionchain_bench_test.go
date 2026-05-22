package versionchain

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/mvcc/versionchain"
)

func BenchmarkVersionChain(
	b *testing.B,
) {

	chain := engine.NewChain()

	value := make([]byte, 128)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		chain.Add(
			uint64(i),
			value,
		)
	}
}
