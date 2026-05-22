package benchmarks

import (
	"bytes"
	"os"
	"testing"

	"github.com/KandakatlaChandramouli/ElasticKV/internal/checkpoint"
)

func BenchmarkMMapCheckpointRestore(
	b *testing.B,
) {

	path := "mmap.snapshot"

	_ = os.Remove(path)

	runtime := checkpoint.NewRuntime()

	payload := bytes.Repeat(
		[]byte("M"),
		1024,
	)

	for i := 0; i < 100000; i++ {

		runtime.Apply(
			uint64(i),
			payload,
		)
	}

	err := runtime.Snapshot(
		path,
	)

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		snapshot, err := checkpoint.LoadMMap(
			path,
		)

		if err != nil {
			b.Fatal(err)
		}

		count := 0

		err = snapshot.Iterate(
			func(
				key uint64,
				value []byte,
			) error {

				count++

				return nil
			},
		)

		if err != nil {
			b.Fatal(err)
		}

		if count != 100000 {

			b.Fatalf(
				"expected 100000 entries got %d",
				count,
			)
		}

		err = snapshot.Close()

		if err != nil {
			b.Fatal(err)
		}
	}
}
