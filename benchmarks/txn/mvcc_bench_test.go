package txn

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/txn"
)

func BenchmarkMVCCRead(
	b *testing.B,
) {

	store := engine.NewStore()

	oracle := engine.NewOracle()

	payload := bytes.Repeat(
		[]byte("T"),
		1024,
	)

	for i := 0; i < 100000; i++ {

		tx := engine.Begin(
			store,
			oracle,
		)

		tx.Put(
			uint64(i),
			payload,
		)
	}

	tx := engine.Begin(
		store,
		oracle,
	)

	target := uint64(77777)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		value, ok := tx.Get(
			target,
		)

		if !ok {
			b.Fatal("missing value")
		}

		if len(value) != 1024 {
			b.Fatal("invalid payload")
		}
	}
}
