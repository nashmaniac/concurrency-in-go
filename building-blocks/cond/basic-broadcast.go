package cond

import (
	"fmt"
	"sync"
	"time"
)

type Button struct {
	cond *sync.Cond
}

func BasicBroadcast() {

	subscribe := func(cond *sync.Cond, fn func()) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			wg.Done()
			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()
			fn()
		}()
		wg.Wait()
	}

	button := Button{
		cond: sync.NewCond(&sync.Mutex{}),
	}
	var mainWg sync.WaitGroup

	mainWg.Add(1)
	subscribe(button.cond, func() {
		defer mainWg.Done()
		time.Sleep(10 * time.Second)
		fmt.Println("Maximizing window")
	})

	mainWg.Add(1)
	subscribe(button.cond, func() {
		defer mainWg.Done()
		time.Sleep(15 * time.Second)
		fmt.Println("Maximizing button")
	})

	mainWg.Add(1)
	subscribe(button.cond, func() {
		defer mainWg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Something is happening")
	})

	button.cond.Broadcast()

	mainWg.Wait()

}
