package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

// GODEBUG=gctrace=1 go run .
// GODEBUG=gctrace=1 GOGC=-1 go run .
func main() {
	debug.SetGCPercent(100)
	debug.SetMemoryLimit(10 * 1024 * 1024) // 10MB
	// memory allocation into memory
	allocateMemory := func(size int) []byte {
		return make([]byte, size)
	}

	// allocating memory to observe its behaviour
	for range 10 {
		allocateMemory(10 * 1024 * 1024)
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v Mib \n", m.Alloc/1024/1024)
		fmt.Printf("TotalAlloc = %v Mib \n", m.TotalAlloc/1024/1024)
		fmt.Printf("Sys = %v Mib \n", m.Sys/1024/1024)
		fmt.Printf("NumGC = %v \n", m.NumGC)
		fmt.Printf("Lookups = %v \n", m.Lookups)
		fmt.Printf("Mallocs = %v \n", m.Mallocs)
		fmt.Printf("Frees = %v \n", m.Frees)
		fmt.Printf("HeapAlloc = %v \n", m.HeapAlloc/1024/1024)
		fmt.Printf("HeapSys = %v \n", m.HeapSys/1024/1024)
		fmt.Printf("HeapIdle = %v \n", m.HeapIdle/1024/1024)
		fmt.Printf("HeapInuse = %v \n", m.HeapInuse/1024/1024)
		fmt.Printf("HeapReleased = %v \n", m.HeapReleased/1024/1024)
		fmt.Printf("HeapObjects = %v \n", m.HeapObjects)
		fmt.Printf("StackInuse = %v \n", m.StackInuse/1024/1024)
		fmt.Printf("StackSys = %v \n", m.StackSys/1024/1024)
		fmt.Printf("MSpanInuse = %v \n", m.MSpanInuse/1024/1024)
		fmt.Printf("MSpanSys = %v \n", m.MSpanSys/1024/1024)
		fmt.Printf("MCacheInuse = %v \n", m.MCacheInuse/1024/1024)
		fmt.Printf("MCacheSys = %v \n", m.MCacheSys/1024/1024)
		fmt.Printf("BuckHashSys = %v \n", m.BuckHashSys/1024/1024)
		fmt.Printf("GCSys = %v \n", m.GCSys/1024/1024)
		fmt.Printf("OtherSys = %v \n", m.OtherSys/1024/1024)
		time.Sleep(1 * time.Second)
		fmt.Println("--------------------------------------------------------")
	}

}
