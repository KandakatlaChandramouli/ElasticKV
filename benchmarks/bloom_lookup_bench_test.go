package benchmarks

import (
	"bytes"
	"os"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/checkpoint"
)

func BenchmarkBloomNegativeLookup(
	b *testing.B,
) {

	path := "bloom.snapshot"

	_ = os.Remove(path)

	runtime := checkpoint.NewRuntime()

	payload := bytes.Repeat([]byte("B"), 1024)

	for i := 0; i < 100000; i++ {
		runtime.Apply(uint64(i), payload)
	}

	err := runtime.Snapshot(path)

	if err != nil {
		b.Fatal(err)
	}

	indexed, err := checkpoint.OpenIndexedSnapshot(path)

	if err != nil {
		b.Fatal(err)
	}

	defer indexed.Close()

	target := uint64(999999999)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_, ok := indexed.Lookup(target)

		if ok {
			b.Fatal("unexpected hit")
		}
	}
}

func BenchmarkBloomPositiveLookup(
	b *testing.B,
) {

	path := "bloom_positive.snapshot"

	_ = os.Remove(path)

	runtime := checkpoint.NewRuntime()

	payload := bytes.Repeat([]byte("P"), 1024)

	for i := 0; i < 100000; i++ {
		runtime.Apply(uint64(i), payload)
	}

	err := runtime.Snapshot(path)

	if err != nil {
		b.Fatal(err)
	}

	indexed, err := checkpoint.OpenIndexedSnapshot(path)

	if err != nil {
		b.Fatal(err)
	}

	defer indexed.Close()

	target := uint64(77777)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		value, ok := indexed.Lookup(target)

		if !ok {
			b.Fatal("missing value")
		}

		if len(value) != 1024 {
			b.Fatal("invalid payload")
		}
	}
}
