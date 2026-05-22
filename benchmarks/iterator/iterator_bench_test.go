package iterator

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/iterator"
)

func BenchmarkIterator(
	b *testing.B,
) {

	data := make(
		[]uint64,
		100000,
	)

	for i := range data {
		data[i] = uint64(i)
	}

	runtime := engine.NewRuntime(
		data,
	)

	var count uint64

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Iterate(
			func(v uint64) {
				count += v
			},
		)
	}

	if count == 0 {
		b.Fatal("iterator failed")
	}
}
