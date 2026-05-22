package lsm

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/lsm"
)

func BenchmarkLSMFlush(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	payload := bytes.Repeat(
		[]byte("L"),
		1024,
	)

	for i := 0; i < 100000; i++ {

		runtime.Put(
			uint64(i),
			payload,
		)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		err := runtime.Flush()

		if err != nil {
			b.Fatal(err)
		}

		for j := 0; j < 100000; j++ {

			runtime.Put(
				uint64(j),
				payload,
			)
		}
	}
}
