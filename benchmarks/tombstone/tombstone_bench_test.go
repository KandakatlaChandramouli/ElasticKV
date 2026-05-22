package tombstone

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/tombstone"
)

func BenchmarkTombstoneLookup(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	for i := 0; i < 100000; i++ {

		runtime.Delete(
			uint64(i),
		)
	}

	target := uint64(77777)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		ok := runtime.IsDeleted(
			target,
		)

		if !ok {
			b.Fatal("lookup failed")
		}
	}
}
