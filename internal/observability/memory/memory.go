package memory

import "runtime"

type Stats struct {
	Alloc uint64
	Sys   uint64
	NumGC uint32
}

func Snapshot() Stats {

	var mem runtime.MemStats

	runtime.ReadMemStats(
		&mem,
	)

	return Stats{
		Alloc: mem.Alloc,
		Sys:   mem.Sys,
		NumGC: mem.NumGC,
	}
}
