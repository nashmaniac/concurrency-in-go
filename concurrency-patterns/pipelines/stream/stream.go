package stream

import "fmt"

func StreamProcessingExample() {
	add := func(num int, adder int) int {
		return num + adder
	}

	multiply := func(num int, multiplier int) int {
		return num * multiplier
	}

	numbers := make([]int, 0)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i+1)
	}

	for _, i := range numbers {
		fmt.Println(add(multiply(i, 2), 1))
	}
}
