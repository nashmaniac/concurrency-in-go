package pool

import (
	"fmt"
	"sync"
)

type Pool struct {
	MemAllocated int
}

func NewPoolObj() Pool {
	return Pool{
		MemAllocated: 1024,
	}
}

func AdvancedPoolTest() {
	var numProcessCreated int = 0

	pool := sync.Pool{
		New: func() any {
			numProcessCreated++
			return NewPoolObj()
		},
	}

	for i := 0; i < 4; i++ {
		pool.Put(pool.New())
	}

	fmt.Printf("%d obj pool worker\n", numProcessCreated)

	var wg sync.WaitGroup
	workCount := 1000
	wg.Add(workCount)
	for i := 0; i < workCount; i++ {
		go func() {
			defer wg.Done()
			p := pool.Get().(Pool)
			defer pool.Put(p)
			fmt.Printf("%p\n", &p)
		}()
	}
	wg.Wait()
	fmt.Printf("%d obj pool worker\n", numProcessCreated)
}
