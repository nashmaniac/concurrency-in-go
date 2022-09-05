package basic

import "fmt"

func generate(done <-chan interface{}, numbers ...int) <-chan int {
	numberStream := make(chan int)

	go func() {
		defer close(numberStream)
		for _, i := range numbers {
			select {
			case <-done:
				return
			case numberStream <- i:
			}
		}
	}()

	return numberStream
}

func add(done <-chan interface{}, numberStream <-chan int, adder int) <-chan int {
	addedStream := make(chan int)

	go func() {
		defer close(addedStream)
		for i := range numberStream {
			select {
			case <-done:
				return
			case addedStream <- i + adder:
			}
		}
	}()

	return addedStream
}

func multiply(done <-chan interface{}, numberStream <-chan int, multiplier int) <-chan int {
	multiplierStream := make(chan int)

	go func() {
		defer close(multiplierStream)
		for i := range numberStream {
			select {
			case <-done:
				return
			case multiplierStream <- i * multiplier:
			}
		}
	}()

	return multiplierStream
}

func BasicPipelineExample() {
	numbers := make([]int, 0)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i+1)
	}

	done := make(chan interface{})
	numberStream := generate(done, numbers...)
	pipeline := add(done, multiply(done, numberStream, 2), 1)

	for i := range pipeline {
		fmt.Println(i)
	}
}
