package main

import (
	"github.com/nashmaniac/concurrency-in-go/locks"
)

func main() {
	// locks.BasicMutex()
	// locks.BasicWithoutMutex()
	locks.BenchMarkMutex()
}
