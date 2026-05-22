package gossip

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/gossip"
)

func BenchmarkGossip(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Broadcast()
	}

	if runtime.Messages.Load() == 0 {
		b.Fatal("broadcast failed")
	}
}
