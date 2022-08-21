package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

/***
should be used when there is no requirements
	* Catching the result
	* there is different mechanism of result  capture
***/
func BasicWaitGroup() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("hello-world-from-first-goroutine")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("hello-world-from-second-goroutine")
	}()

	wg.Wait()
}

func printNumber(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	fmt.Println(time.Now(), num)
}

func WaitGroupWithPassingReferences() {
	var wg sync.WaitGroup

	numProcess := 5
	for i := 0; i < numProcess; i++ {
		wg.Add(1)
		go printNumber(&wg, i+1)
	}
	wg.Wait()

}
