package admission

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/admission"
)

func BenchmarkAdmission(
	b *testing.B,
) {

	runtime := engine.NewRuntime(
		uint64(b.N + 1),
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		ok := runtime.Allow(
			uint64(i),
		)

		if !ok {
			b.Fatal("admission failed")
		}
	}
}
