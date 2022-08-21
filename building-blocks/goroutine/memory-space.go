package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

func TestMemorySpaceConsumption() {
	memConsumed := func() uint64 {
		runtime.GC()
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		return mem.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup

	noop := func() {
		wg.Done()
		<-c
	}

	before := memConsumed()
	routineCount := 1e5
	wg.Add(int(routineCount))
	for i := 0; i < int(routineCount); i++ {
		go noop()
	}
	wg.Wait()
	after := memConsumed()

	fmt.Printf("%.2f kb\n", float64(after-before)/routineCount/1000)

}
