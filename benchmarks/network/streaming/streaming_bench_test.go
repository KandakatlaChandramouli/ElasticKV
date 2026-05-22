package streaming

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/network/streaming"
)

func BenchmarkStreaming(
	b *testing.B,
) {

	stream := engine.NewStream()

	payload := make([]byte, 4096)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		stream.Send(
			payload,
		)
	}
}
