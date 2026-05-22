package leader

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/leader"
)

func BenchmarkLeaderElection(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Elect(
			7,
		)

		current := runtime.Current()

		if current != 7 {
			b.Fatal("election failed")
		}
	}
}
