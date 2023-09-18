package main

import (
	"fmt"
	"runtime"
)

func main() {
	var memStats runtime.MemStats
	//s := `⓴`
	runtime.ReadMemStats(&memStats)
	fmt.Println("Allocated memory:", memStats.Alloc)
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)
	a := "a"
	runtime.ReadMemStats(&memStats)
	fmt.Println("Allocated memory:", memStats.Alloc)
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)
	b := "b"
	runtime.ReadMemStats(&memStats)
	fmt.Println("Allocated memory:", memStats.Alloc)
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)
	c := "c"
	d := "d"
	e := "e"
	f := "f"
	g := "g"

	str := a + b + c + d + e + f + g
	fmt.Println(str)
	runtime.GC()
	runtime.ReadMemStats(&memStats)
	fmt.Println("Allocated memory:", memStats.Alloc)
	fmt.Println("HeapAlloc memory:", memStats.HeapAlloc)
}
