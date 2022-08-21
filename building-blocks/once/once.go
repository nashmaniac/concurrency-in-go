package once

import (
	"fmt"
	"sync"
)

func OnceTest() {

	var count, count1, a, b int

	increment := func() {
		count++
		a--
	}
	decrement := func() {
		count1--
		b += 5
	}

	var oneTime sync.Once
	oneTime.Do(decrement)
	oneTime.Do(increment)

	fmt.Println(count, count1)
	fmt.Println(a, b)

}
