package crashsim

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/crashsim"
)

func BenchmarkCrashSimulation(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		runtime.Crash()
	}

	if runtime.Crashes.Load() == 0 {
		b.Fatal("crash failed")
	}
}
