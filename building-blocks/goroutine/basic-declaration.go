package goroutine

import (
	"fmt"
	"sync"
)

func hello_world() {
	fmt.Println("hello-world")
}

func WithFunction() {
	go hello_world()
	// time.Sleep(5 * time.Second)
}

func InlineFunction() {
	hello_world := func() {
		fmt.Println("hello-world-from-inline-function")
	}

	go hello_world()
	// time.Sleep(5 * time.Second) // should avoid as this would not gurantee the execution
}

func AnonymousFunction() {
	go func() {
		fmt.Println("hello-world-from-anonymous-function")
	}()
	// time.Sleep(5 * time.Second)
}

func WithSyncGroup() {
	var wg sync.WaitGroup // declaring the wait model

	wg.Add(1) // number of go routine to wait

	go func() {
		defer wg.Done() // acknowledge end of the goroutine
		fmt.Println("hello-world-from-fork-join-model")
	}()

	wg.Wait() // waiting from the main goroutine for finish
}
