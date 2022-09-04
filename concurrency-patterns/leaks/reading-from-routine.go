package leaks

import (
	"fmt"
	"time"
)

func ReadingFromRoutine() {
	// write the reader

	consumer := func(ch <-chan int, done <-chan interface{}) <-chan interface{} {
		terminated := make(chan interface{})

		go func() {
			defer close(terminated)
			defer fmt.Println("returning from the routine")

			for {
				select {
				case c := <-ch:
					fmt.Printf("reading %d\n", c)
				case <-done:
					return
				}
			}

		}()

		return terminated
	}

	producer := func(ch chan<- int, done chan<- interface{}) {
		go func() {
			for i := 0; i < 10; i++ {
				ch <- i + 1
				time.Sleep(500 * time.Millisecond)
			}
			close(done)
		}()

	}

	// write the producer
	ch := make(chan int)
	done := make(chan interface{})

	terminated := consumer(ch, done)
	producer(ch, done)
	<-terminated
	fmt.Println("done with the routine fully")

}
