package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)

	fmt.Printf("stats: %d\n", stats.HeapAlloc)

	bigArray := [1 << 25]int{}

	runtime.ReadMemStats(stats)
	fmt.Printf("stats after bigArray created: %d\n", stats.HeapAlloc)
	_ = bigArray

	//go funcWithArray(bigArray)
	//
	//<-time.After(time.Second * 1)
	//runtime.ReadMemStats(stats)
	//fmt.Printf("stats bigArray copied to gourutine stack: %d\n", stats.HeapAlloc)

	go funcWithArrayPointer(&bigArray)

	<-time.After(time.Second * 1)
	runtime.ReadMemStats(stats)
	fmt.Printf("stats bigArray copied to gourutine stack: %d\n", stats.HeapAlloc)

	time.Sleep(time.Hour)
	_ = bigArray
}

func funcWithArray(a [1 << 25]int) {
	time.Sleep(time.Hour)
	_ = a
}

func funcWithArrayPointer(a *[1 << 25]int) {
	time.Sleep(time.Hour)
	_ = a
}
