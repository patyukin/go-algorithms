package main

import (
	"fmt"
	"runtime"
)

func main() {
	var memStats runtime.MemStats
	m := make(map[int]int64)
	runtime.ReadMemStats(&memStats)
	for i := 0; i < 1_000_000; i++ {
		m[i] = 9949999254234523
	}

	fmt.Println(len(m))
	fmt.Println("Allocated memory:", memStats.Alloc)
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)

	clear(m)
	runtime.ReadMemStats(&memStats)
	fmt.Println(len(m))
	fmt.Println("Allocated memory:", memStats.Alloc)
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)

	for i := 0; i < 1_000_000; i++ {
		m[i] = 9949999254234523
	}
	runtime.GC()
	runtime.ReadMemStats(&memStats)
	fmt.Println(len(m))
	fmt.Println("Allocated memory:", memStats.Alloc)
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)

	for i := 0; i < 1_000_000; i++ {
		m[i] = 9949999254234523
	}
	m = make(map[int]int64)
	runtime.ReadMemStats(&memStats)
	fmt.Println(len(m))
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)
	fmt.Println("Allocated memory:", memStats.Alloc)

	for i := 0; i < 1_000_000; i++ {
		m[i] = 9949999254234523
	}
	m = nil
	runtime.ReadMemStats(&memStats)
	fmt.Println(len(m))
	fmt.Println("Allocated memory:", memStats.Alloc)
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)
}
