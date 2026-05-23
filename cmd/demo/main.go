package main

import (
	"math/rand"

	executor "github.com/KandakatlaChandramouli/ElasticKV/internal/pipeline/executor"

	ctx "github.com/KandakatlaChandramouli/ElasticKV/internal/runtime/context"
)

func main() {

	vector := make(
		[]float32,
		768,
	)

	for i := range vector {

		vector[i] = rand.Float32()
	}

	query := ctx.QueryContext{
		Query: "semantic vector retrieval",
		Vector: vector,
		TopK: 10,
		EnableGPU: true,
		Distributed: true,
	}

	executor.Execute(
		query,
	)
}
