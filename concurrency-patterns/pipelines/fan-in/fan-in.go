package fanin

import (
	"fmt"
	"sync"
	"time"
)

func generator(done <-chan interface{}, n int) <-chan interface{} {
	numberStream := make(chan interface{})

	go func() {
		defer close(numberStream)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case numberStream <- i + 1:
			}
		}
	}()

	return numberStream
}

func delayFeed(done <-chan interface{}, delay time.Duration, numberStream <-chan interface{}) <-chan interface{} {
	feedStream := make(chan interface{})

	go func() {
		defer close(feedStream)

		for i := range numberStream {
			time.Sleep(delay)
			select {
			case <-done:
				return
			case feedStream <- i:
			}
		}
	}()

	return feedStream
}

func fanIn(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	multiplexedStream := make(chan interface{})
	var wg sync.WaitGroup
	multiplex := func(wg *sync.WaitGroup, c <-chan interface{}) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- i:
			}
		}
	}

	n := len(channels)
	wg.Add(n)

	for _, c := range channels {
		go multiplex(&wg, c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func BasicExampleWithOutFanInFanOut() {
	n := 10
	delay := 2 * time.Second

	done := make(chan interface{})
	defer close(done)

	s := time.Now()
	numberStream := generator(done, n)
	pipeline := delayFeed(done, delay, numberStream)
	for i := range pipeline {
		fmt.Println(i)
	}
	fmt.Printf("time taken: %f secs", time.Since(s).Seconds())
}

func AdvancedExample() {
	n := 32
	delay := 1 * time.Second
	numCpu := 32
	fanOuts := make([]<-chan interface{}, numCpu)
	done := make(chan interface{})
	defer close(done)

	s := time.Now()
	numberStream := generator(done, n)
	for i := 0; i < numCpu; i++ {
		fanOuts[i] = delayFeed(done, delay, numberStream)
	}
	pipeline := fanIn(done, fanOuts...)
	for i := range pipeline {
		fmt.Println(i)
	}
	fmt.Printf("time taken: %f secs", time.Since(s).Seconds())
}
