package operators

import (
	"testing"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/queryexecutor/operators"
)

func BenchmarkOperators(
	b *testing.B,
) {

	runtime := engine.NewRuntime()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		if !runtime.Execute() {
			b.Fatal("execution failed")
		}
	}
}
