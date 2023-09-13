package main

import (
	"fmt"
	"runtime"
)

func main() {
	var memStats runtime.MemStats
	m := make([]int64, 0, 1_000_000)
	runtime.ReadMemStats(&memStats)
	for i := 0; i < 1_000_000; i++ {
		m = append(m, 9949999254234523)
	}

	fmt.Println("len", len(m))
	fmt.Println("cap", cap(m))
	fmt.Println("Allocated memory:", memStats.Alloc)
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)
	fmt.Println(m[0])
	clear(m)
	runtime.ReadMemStats(&memStats)
	fmt.Println("after clear: [0] = ", m[0])
	fmt.Println("after clear: len = ", len(m))
	fmt.Println("after clear: cap = ", cap(m))
	fmt.Println("Allocated memory:", memStats.Alloc)
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)

	for i := 0; i < 1_000_000; i++ {
		m = append(m, 9949999254234523)
	}
	runtime.GC()
	runtime.ReadMemStats(&memStats)
	fmt.Println("after CG: len = ", len(m))
	fmt.Println("after CG: cap = ", cap(m))
	fmt.Println("Allocated memory:", memStats.Alloc)
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)

	for i := 0; i < 1_000_000; i++ {
		m = append(m, 9949999254234523)
	}
	m = make([]int64, 0)
	runtime.ReadMemStats(&memStats)
	fmt.Println("after make: len = ", len(m))
	fmt.Println("after make: cap = ", cap(m))
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)
	fmt.Println("Allocated memory:", memStats.Alloc)

	for i := 0; i < 1_000_000; i++ {
		m = append(m, 9949999254234523)
	}
	m = nil
	runtime.ReadMemStats(&memStats)
	fmt.Println("after nil: len = ", len(m))
	fmt.Println("after nil: cap = ", cap(m))
	fmt.Println("Allocated memory:", memStats.Alloc)
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)
}
