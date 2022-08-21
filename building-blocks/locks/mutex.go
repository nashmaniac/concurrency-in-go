package locks

import (
	"fmt"
	"sync"
)

func BasicWithoutMutex() {
	var count int64 = 0
	numCount := 50

	increment := func() {
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}
	decrement := func() {
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithmatic sync.WaitGroup
	arithmatic.Add(numCount)
	for i := 0; i < numCount; i++ {
		go func() {
			defer arithmatic.Done()
			increment()
		}()
	}
	arithmatic.Add(numCount)
	for i := 0; i < numCount; i++ {
		go func() {
			defer arithmatic.Done()
			decrement()
		}()
	}

	arithmatic.Wait()

}

func BasicMutex() {
	var count int64 = 0
	numCount := 50

	var c sync.Mutex
	increment := func() {
		c.Lock()
		defer c.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}
	decrement := func() {
		c.Lock()
		defer c.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithmatic sync.WaitGroup
	arithmatic.Add(numCount)
	for i := 0; i < numCount; i++ {
		go func() {
			defer arithmatic.Done()
			increment()
		}()
	}
	arithmatic.Add(numCount)
	for i := 0; i < numCount; i++ {
		go func() {
			defer arithmatic.Done()
			decrement()
		}()
	}

	arithmatic.Wait()
}
