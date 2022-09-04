package forselect

import (
	"fmt"
	"time"
)

func ForSelectExample() {

	ch := make(chan int)
	done := make(chan interface{})

	go func() {
		defer fmt.Println("return from the goroutine")
		for {
			select {
			case c := <-ch:
				fmt.Println(c)
			case <-done:
				return
				// fmt.Print("-")
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(1 * time.Second)
	}
	close(done)
	time.Sleep(1 * time.Second)
}
