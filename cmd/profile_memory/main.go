package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	engine "github.com/KandakatlaChandramouli/ElasticKV/internal/hnsw/core"
)

func main() {

	f, _ := os.Create("cpu.prof")

	pprof.StartCPUProfile(f)

	defer pprof.StopCPUProfile()

	start := time.Now()

	index := engine.New()

	vector := make([]float32, 768)

	for i := 0; i < 5000; i++ {

		index.Insert(
			i,
			vector,
		)
	}

	index.Search(
		vector,
		10,
	)

	fmt.Println(
		"elapsed:",
		time.Since(start),
	)
}
