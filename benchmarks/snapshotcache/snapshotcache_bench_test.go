package snapshotcache

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/snapshotcache"
)

func BenchmarkSnapshotCache(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Put(uint64(i))
	}

	if len(runtime.Cache) == 0 {
		b.Fatal("snapshot cache failed")
	}
}
