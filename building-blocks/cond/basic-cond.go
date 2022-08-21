package cond

import (
	"fmt"
	"sync"
	"time"
)

func BasicCond() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	count := 0
	printAndSignal := func(waitTime int, input int) {
		time.Sleep(time.Duration(waitTime) * time.Second)
		cond.L.Lock()
		defer wg.Done()
		defer cond.L.Unlock()
		fmt.Printf("Printing %d\n", input)
		count--
		cond.Signal()
	}

	for i := 0; i < 10; i++ {
		cond.L.Lock()
		for count == 2 {
			cond.Wait()
		}
		cond.L.Unlock()
		fmt.Printf("Pushing %d\n", i+1)
		count++
		wg.Add(1)
		go printAndSignal(1, i+1)
	}
	wg.Wait()
}
