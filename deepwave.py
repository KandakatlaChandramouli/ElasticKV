from pathlib import Path

files = {

"internal/vectorkernel/cosine/runtime.go": r'''package cosine

func Similarity(
    a []float32,
    b []float32,
) float32 {

    var dot float32
    var normA float32
    var normB float32

    for i := 0; i < len(a); i++ {

        dot += a[i] * b[i]

        normA += a[i] * a[i]

        normB += b[i] * b[i]
    }

    if normA == 0 || normB == 0 {
        return 0
    }

    return dot
}
''',

"internal/vectorkernel/dotproduct/runtime.go": r'''package dotproduct

func Compute(
    a []float32,
    b []float32,
) float32 {

    var result float32

    for i := 0; i < len(a); i++ {
        result += a[i] * b[i]
    }

    return result
}
''',

"internal/vectorkernel/l2/runtime.go": r'''package l2

func Distance(
    a []float32,
    b []float32,
) float32 {

    var distance float32

    for i := 0; i < len(a); i++ {

        diff := a[i] - b[i]

        distance += diff * diff
    }

    return distance
}
''',

"internal/vectorkernel/topk/runtime.go": r'''package topk

import "sort"

func Select(
    values []float32,
    k int,
) []float32 {

    sorted := make(
        []float32,
        len(values),
    )

    copy(
        sorted,
        values,
    )

    sort.Slice(
        sorted,
        func(i, j int) bool {
            return sorted[i] > sorted[j]
        },
    )

    return sorted[:k]
}
''',

"benchmarks/vectorkernel/cosine/cosine_bench_test.go": r'''package cosine

import (
    "testing"

    engine "github.com/KandakatlaChandramouli/ElasticKV/internal/vectorkernel/cosine"
)

func BenchmarkCosine(
    b *testing.B,
) {

    a := make([]float32, 1024)
    c := make([]float32, 1024)

    for i := 0; i < 1024; i++ {
        a[i] = float32(i)
        c[i] = float32(i)
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        engine.Similarity(a, c)
    }
}
''',

"benchmarks/vectorkernel/dotproduct/dotproduct_bench_test.go": r'''package dotproduct

import (
    "testing"

    engine "github.com/KandakatlaChandramouli/ElasticKV/internal/vectorkernel/dotproduct"
)

func BenchmarkDotProduct(
    b *testing.B,
) {

    a := make([]float32, 1024)
    c := make([]float32, 1024)

    for i := 0; i < 1024; i++ {
        a[i] = float32(i)
        c[i] = float32(i)
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        engine.Compute(a, c)
    }
}
''',

"benchmarks/vectorkernel/l2/l2_bench_test.go": r'''package l2

import (
    "testing"

    engine "github.com/KandakatlaChandramouli/ElasticKV/internal/vectorkernel/l2"
)

func BenchmarkL2(
    b *testing.B,
) {

    a := make([]float32, 1024)
    c := make([]float32, 1024)

    for i := 0; i < 1024; i++ {
        a[i] = float32(i)
        c[i] = float32(i)
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        engine.Distance(a, c)
    }
}
''',

"benchmarks/vectorkernel/topk/topk_bench_test.go": r'''package topk

import (
    "testing"

    engine "github.com/KandakatlaChandramouli/ElasticKV/internal/vectorkernel/topk"
)

func BenchmarkTopK(
    b *testing.B,
) {

    values := make([]float32, 10000)

    for i := 0; i < 10000; i++ {
        values[i] = float32(i)
    }

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        engine.Select(values, 10)
    }
}
''',
}

for path, content in files.items():

    Path(path).parent.mkdir(
        parents=True,
        exist_ok=True,
    )

    with open(path, "w") as f:
        f.write(content)

print("DEEP COMPUTE WAVE READY")
