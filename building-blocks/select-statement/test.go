package selectstatement

import (
	"fmt"
	"sync"
)

func TestWork() {
	var wg sync.WaitGroup
	wg.Add(2)

	num := 1000
	ch := make(chan int)
	go func(n int, wg *sync.WaitGroup, ch chan<- int) {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- i + 1
			fmt.Println("publishing ", i+1)
			// time.Sleep(1 * time.Second)
		}
	}(num, &wg, ch)

	go func(wg *sync.WaitGroup, ch <-chan int) {
		defer wg.Done()
		for i := range ch {
			fmt.Println("Received ", i)
			// time.Sleep(2 * time.Second)
		}
	}(&wg, ch)

	wg.Wait()
}

func TestWork1() {

	producer := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 50; i++ {
				ch <- i + 1
			}
		}()
		return ch
	}

	consumer := func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
	}

	ch := producer()
	consumer(ch)
}
