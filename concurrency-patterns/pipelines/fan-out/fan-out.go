package fanout

import (
	"fmt"
	"runtime"
	"time"
)

func generator(done <-chan interface{}, values ...int) <-chan interface{} {
	valueStream := make(chan interface{})

	go func() {
		defer close(valueStream)

		for _, i := range values {
			select {
			case <-done:
				return
			case valueStream <- i:
			}
		}
	}()

	return valueStream
}

func feedNumber(done <-chan interface{}, valueStream <-chan interface{}) <-chan interface{} {
	feedStream := make(chan interface{})

	go func() {
		defer close(feedStream)

		for i := range valueStream {
			time.Sleep(1 * time.Second)
			select {
			case <-done:
				return
			case feedStream <- i:
			}
		}

	}()

	return feedStream
}

func WithoutFanoutBasicExample() {
	done := make(chan interface{})
	defer close(done)
	numbers := make([]int, 0)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i+1)
	}

	valueStream := generator(done, numbers...)
	pipeline := feedNumber(done, valueStream)

	s := time.Now()
	for i := range pipeline {
		fmt.Println(i)
	}
	fmt.Printf("Time taken: %f\n", time.Since(s).Seconds())
}

func FanOutExample() {
	numCPU := runtime.NumCPU()
	fanOut := make([]<-chan interface{}, numCPU)

	done := make(chan interface{})
	defer close(done)
	numbers := make([]int, 0)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i+1)
	}

	valueStream := generator(done, numbers...)
	for i := 0; i < numCPU; i++ {
		fanOut[i] = feedNumber(done, valueStream)
	}

	s := time.Now()
	fmt.Printf("Time taken: %f\n", time.Since(s).Seconds())

}
