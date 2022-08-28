package main

import (
	selectstatement "github.com/nashmaniac/concurrency-in-go/building-blocks/select-statement"
)

func main() {
	// locks.BasicMutex()
	// locks.BasicWithoutMutex()
	// cond.BasicCond()
	// cond.BasicBroadcast()
	// once.OnceTest()

	// pool.AdvancedPoolTest()
	// channels.BufferedChannel()
	// selectstatement.BasicChannelTest()
	selectstatement.TestWork1()
}
