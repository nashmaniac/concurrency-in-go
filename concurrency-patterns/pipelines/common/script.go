package common

import (
	"fmt"
	"math/rand"
)

func RepeatPipeline() {
	done := make(chan interface{})
	defer close(done)
	pipeline := TakeStage(done, RepeatStage(done, 5), 15)

	for i := range pipeline {
		fmt.Println(i)
	}
}

func RepeatFunctiontPipeline() {
	done := make(chan interface{})
	defer close(done)
	random := func() interface{} {
		return rand.Int()
	}
	pipeline := TakeStage(done, RepeatFunctionStage(done, random), 15)

	for i := range pipeline {
		fmt.Println(i)
	}
}
