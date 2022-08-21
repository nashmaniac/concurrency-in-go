package goroutine

import (
	"sync"
	"testing"
)

func BenchMarkContextSwitching(b *testing.B) {

	var wg sync.WaitGroup
	var c chan interface{}
	var begin chan interface{}
	c = make(chan interface{})
	begin = make(chan interface{})

	var token struct{}
	producer := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	consumer := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go producer()
	go consumer()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
