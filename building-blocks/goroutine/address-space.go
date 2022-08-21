package goroutine

import (
	"fmt"
	"sync"
)

func ClosureVariableTest() {
	// test variable changes across the goroutine
	sampleInput := "hello"
	fmt.Printf("before: %s\n", sampleInput)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		sampleInput = "world"
	}()

	wg.Wait()
	fmt.Printf("after: %s\n", sampleInput)
}

func LoopVariableTestWithSameAddressSpace() {
	inputList := []string{
		"hello", "world", "shetu",
	}

	var wg sync.WaitGroup
	wg.Add(len(inputList))

	for _, word := range inputList {
		go func() {
			defer wg.Done()
			fmt.Printf("%s\n", word)
		}()
	}

	wg.Wait()
}

func LoopVariableWithDifferentAddressSpace() {
	inputList := []string{
		"hello", "world", "shetu",
	}

	var wg sync.WaitGroup
	wg.Add(len(inputList))

	for _, word := range inputList {
		go func(word string) {
			defer wg.Done()
			fmt.Printf("%s\n", word)
		}(word)
	}

	wg.Wait()
}
