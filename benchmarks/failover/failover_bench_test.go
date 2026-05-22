package failover

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/failover"
)

func BenchmarkFailover(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Switch()
	}

	if runtime.Switches.Load() == 0 {
		b.Fatal("failover failed")
	}
}
