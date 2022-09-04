package leaks

import (
	"fmt"
	"math/rand"
	"time"
)

func WritingFromRoutine() {

	// write a producer that infinitely write to a data stream

	producer := func(ch chan<- int, done <-chan interface{}) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer close(terminated)
			defer fmt.Println("getting out from producer")
			for {
				select {
				case ch <- rand.Int():
				case <-done:
					fmt.Println("got signal to stop")
					return
				}
			}
		}()
		return terminated
	}

	// write a consumer that only consumes given amount of number and notify to stop
	consumer := func(n int, ch <-chan int, done chan<- interface{}) {
		go func() {
			defer close(done)
			defer fmt.Println("getting out from consumer")
			defer fmt.Println("stopping request")
			fmt.Println("will read ", n, " numbers")
			for i := 0; i < n; i++ {
				fmt.Printf("reading %d\n", <-ch)
				time.Sleep(1 * time.Second)
			}
		}()
	}

	// simulate it

	ch := make(chan int)
	done := make(chan interface{})
	n := 10
	terminated := producer(ch, done)
	consumer(n, ch, done)
	<-terminated
	fmt.Println("done fully")
}
