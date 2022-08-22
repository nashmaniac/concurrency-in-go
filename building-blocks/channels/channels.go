package channels

import (
	"fmt"
	"sync"
	"time"
)

func BlockingWithChannel() {
	channel := make(chan interface{})
	go func() {
		channel <- struct{}{}
	}()
	fmt.Println(<-channel)
}

func DuplexChannel() {
	channel := make(chan string)
	go func() {
		channel <- "hello world"
	}()

	fmt.Println(<-channel)
}

func ReadOnClosedChannel() {
	channel := make(chan string)
	close(channel)
	input, ok := <-channel
	fmt.Println(input, ok)
	// close(channel)
	// input, ok = <-channel
	// fmt.Println(input, ok)
}

func RangingOverChannel() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			fmt.Printf("Inputting %d at %s\n", i+1, time.Now())
			channel <- i + 1
		}
	}()

	for i := range channel {
		fmt.Printf("Removing %d at %s\n", i, time.Now())
	}
}

func TriggerChannelByClosing() {
	numProcess := 10

	var wg sync.WaitGroup
	wg.Add(numProcess)
	begin := make(chan bool)
	for i := 0; i < numProcess; i++ {
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%d has begun %s\n", i, time.Now())
		}(i)
	}

	close(begin)
	wg.Wait()
}

func BufferedChannel() {
	channels := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(channels)
		defer fmt.Println("producer done")
		for i := 0; i < 10; i++ {
			channels <- i + 1
		}
	}()
	for i := range channels {
		fmt.Println(i)
	}
	wg.Wait()
}
