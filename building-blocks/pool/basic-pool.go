package pool

import (
	"fmt"
	"sync"
)

func BasicPoolTest() {
	pool := &sync.Pool{
		New: func() any {
			fmt.Println("new pool is getting created")
			return &struct{}{}
		},
	}

	myPool := pool.Get()
	defer pool.Put(myPool)
	secondPool := pool.Get()
	defer pool.Put(secondPool)
}

type PoolObject struct {
	count int
}

func NewPoolObject() PoolObject {
	return PoolObject{
		count: 5,
	}
}

func (p *PoolObject) Reset() PoolObject {
	p.count = 5
	return *p
}

func (p *PoolObject) GetCount() int {
	return p.count
}

func (p *PoolObject) SetCount(count int) {
	p.count = count
}

func BasicPoolTestWithReset() {
	pool := sync.Pool{
		New: func() any {
			fmt.Println("creating new object")
			return NewPoolObject()
		},
	}

	p := pool.Get().(PoolObject)
	p1 := pool.Get().(PoolObject)
	defer pool.Put(p.Reset())
	defer pool.Put(p1.Reset())

	fmt.Println(p.GetCount())
	p.SetCount(1)
	fmt.Println(p.GetCount())
	fmt.Println(p1.GetCount())
	p1.SetCount(10)
	fmt.Println(p1.GetCount())

}
