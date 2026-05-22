package sstable

import (
	"bytes"
	"os"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/sstable"
)

func BenchmarkSSTableLookup(
	b *testing.B,
) {

	path := "table.sst"

	_ = os.Remove(path)

	entries := make(
		[]engine.Entry,
		0,
		100000,
	)

	payload := bytes.Repeat(
		[]byte("S"),
		1024,
	)

	for i := 0; i < 100000; i++ {

		entries = append(
			entries,
			engine.Entry{
				Key:   uint64(i),
				Value: payload,
			},
		)
	}

	err := engine.Build(
		path,
		entries,
	)

	if err != nil {
		b.Fatal(err)
	}

	table, err := engine.Open(path)

	if err != nil {
		b.Fatal(err)
	}

	defer table.Close()

	target := uint64(77777)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		value, ok := table.Lookup(
			target,
		)

		if !ok {
			b.Fatal("lookup failed")
		}

		if len(value) != 1024 {
			b.Fatal("invalid payload")
		}
	}
}
