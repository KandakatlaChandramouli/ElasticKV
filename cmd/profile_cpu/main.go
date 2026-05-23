package main

import (
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
)

func main() {

	file, err := os.Create(
		"cpu.prof",
	)

	if err != nil {

		log.Fatal(err)
	}

	pprof.StartCPUProfile(
		file,
	)

	defer pprof.StopCPUProfile()

	buffer := make(
		[]float64,
		10_000_000,
	)

	for i := range buffer {

		buffer[i] = rand.Float64() * rand.Float64()
	}
}
