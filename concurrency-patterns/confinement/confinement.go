package confinement

import "fmt"

func AdhocExample() {
	data := make([]int, 0)
	for i := 0; i < 10; i++ {
		data = append(data, i+1)
	}

	dataStream := make(chan int)
	go func() {
		defer close(dataStream)
		for _, i := range data {
			dataStream <- i
		}
	}()

	for i := range dataStream {
		fmt.Println(i)
	}
}

func LexicalExample() {
	doWork := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 10; i++ {
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
	ch := doWork()
	consumer(ch)
}
