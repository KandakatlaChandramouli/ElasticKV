package logbuffer

import (
	"bytes"
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/logbuffer"
)

func BenchmarkLogBuffer(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	payload := bytes.Repeat(
		[]byte("L"),
		1024,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		runtime.Append(payload)
	}

	if len(runtime.Entries) == 0 {
		b.Fatal("append failed")
	}
}
