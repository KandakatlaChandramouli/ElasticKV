package main

import (
    "fmt"
    "time"

    engine "github.com/KandakatlaChandramouli/ElasticKV/internal/gpu/batchdistance"
)

func main() {

    vectors := make([][]float32, 100000)

    for i := range vectors {
        vectors[i] = make([]float32, 768)
    }

    query := make([]float32, 768)

    start := time.Now()

    scores := engine.Compute(
        vectors,
        query,
    )

    elapsed := time.Since(start)

    fmt.Println(
        "vectors:",
        len(scores),
    )

    fmt.Println(
        "elapsed:",
        elapsed,
    )
}
