package ringbuffer

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/ringbuffer"
)

func BenchmarkRingBufferPush(
	b *testing.B,
) {

	runtime := engine.NewRuntime(
		1024,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Push(
			uint64(i),
		)
	}

	if runtime.Head == 0 {
		b.Fatal("push failed")
	}
}
