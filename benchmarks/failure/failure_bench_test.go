package failure

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/failure"
)

func BenchmarkFailureInjection(
	b *testing.B,
) {

	injector := engine.NewInjector(
		1.0,
	)

	var failures uint64

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		if injector.Fail() {
			failures++
		}
	}

	if failures == 0 {
		b.Fatal("injector inactive")
	}
}
