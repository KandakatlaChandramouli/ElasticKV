package vectorcache

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/vectorcache"
)

func BenchmarkVectorCache(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Put(uint64(i))
	}

	if len(runtime.Cache) == 0 {
		b.Fatal("vector cache failed")
	}
}
