package main

import (
	"fmt"
	"runtime"
)

func main() {

	var mem runtime.MemStats

	runtime.ReadMemStats(
		&mem,
	)

	fmt.Println(
		"\n=== ElasticKV Memory Profile ===\n",
	)

	fmt.Printf(
		"Alloc = %v MB\n",
		mem.Alloc / 1024 / 1024,
	)

	fmt.Printf(
		"TotalAlloc = %v MB\n",
		mem.TotalAlloc / 1024 / 1024,
	)

	fmt.Printf(
		"Sys = %v MB\n",
		mem.Sys / 1024 / 1024,
	)

	fmt.Printf(
		"NumGC = %v\n",
		mem.NumGC,
	)
}
