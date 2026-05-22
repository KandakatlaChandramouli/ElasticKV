package memtable

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/lsm/memtable"
)

func BenchmarkMemTable(
	b *testing.B,
) {

	table := engine.New()

	value := make([]byte, 256)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		table.Put(
			"key",
			value,
		)
	}
}
