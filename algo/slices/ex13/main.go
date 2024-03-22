package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	b := make([]int, 0, 1_000_000)

	runtime.ReadMemStats(&m)
	getMem(&m)
	//var b []int
	for i := 0; i < 1_000_000; i++ {
		b = append(b, i)
	}

	runtime.GC()
	runtime.ReadMemStats(&m)
	getMem(&m)
}

func bToMb(b uint64) uint64 {
	return b / 1024
}
func getMem(m *runtime.MemStats) {
	fmt.Printf(
		"4. Alloc = %v KiB\tTotalAlloc = %v KiB\tSys = %v KiB\tNumGC = %v\n",
		bToMb(m.Alloc),
		bToMb(m.TotalAlloc),
		bToMb(m.Sys), m.NumGC,
	)
}
