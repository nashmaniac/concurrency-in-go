package main

import (
	"github.com/nashmaniac/concurrency-in-go/building-blocks/once"
)

func main() {
	// locks.BasicMutex()
	// locks.BasicWithoutMutex()
	// cond.BasicCond()
	// cond.BasicBroadcast()
	once.OnceTest()
}
