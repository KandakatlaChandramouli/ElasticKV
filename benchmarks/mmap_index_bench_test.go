package benchmarks

import (
	"bytes"
	"os"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/checkpoint"
)

func BenchmarkMMapIndexedLookup(
	b *testing.B,
) {

	path := "indexed.snapshot"

	_ = os.Remove(path)

	runtime := checkpoint.NewRuntime()

	payload := bytes.Repeat([]byte("I"), 1024)

	for i := 0; i < 100000; i++ {
		runtime.Apply(uint64(i), payload)
	}

	err := runtime.Snapshot(path)

	if err != nil {
		b.Fatal(err)
	}

	snapshot, err := checkpoint.LoadMMap(path)

	if err != nil {
		b.Fatal(err)
	}

	defer snapshot.Close()

	index := checkpoint.BuildIndex(snapshot.Data)

	target := uint64(77777)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		entry, ok := index.Lookup(target)

		if !ok {
			b.Fatal("lookup failed")
		}

		value := snapshot.Data[entry.Offset : entry.Offset+entry.Length]

		if len(value) != 1024 {
			b.Fatal("invalid payload")
		}
	}
}
