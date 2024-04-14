package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

type row struct {
	key       string
	val       int
	ttl       time.Duration
	createdAt time.Time
}

func main() {
	if len(os.Args) < 2 {
		panic("please choose version to run i.e. 'go run . v1'")
	}
	version := os.Args[1]

	cpu, _ := os.Create(version + "cpu.prof")
	if err := pprof.StartCPUProfile(cpu); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	mem, _ := os.Create(version + "mem.prof")
	runtime.GC()

	fmt.Printf("executing %v\n", version)
	switch version {
	case "v1":
		v1()
	}

	if err := pprof.WriteHeapProfile(mem); err != nil {
		panic(err)
	}
}
