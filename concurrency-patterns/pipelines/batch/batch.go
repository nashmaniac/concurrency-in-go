package batch

import "fmt"

// a stage of the pipeline should input and output the same thing
// a stage should not modify the initial input data
func multiply(numbers []int, mulitiplier int) []int {
	modifiedNumber := make([]int, len(numbers))
	for i := range numbers {
		modifiedNumber[i] = numbers[i] * mulitiplier
	}
	return modifiedNumber
}

func add(numbers []int, adder int) []int {
	addedNumber := make([]int, len(numbers))
	for i := range numbers {
		addedNumber[i] = numbers[i] + adder
	}
	return addedNumber
}

func BatchProcessingExample() {
	numbers := make([]int, 0)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i+1)
	}

	for _, i := range add(multiply(numbers, 2), 1) {
		fmt.Println(i)
	}
}

// bottleneck
// 1. memory footprint is double
