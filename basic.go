package main

import (
	"fmt"
)

/**
Fancy way of doing: [0, 1, 6, 388, 12387].map(v => v*2).map(v => v+3)
**/
func main() {
	// batch of numbers
	nums := []int{0, 1, 6, 388, 12387}

	// create input channel (source) and send numbers
	inputChan := make(chan int, len(nums))
	go func() {
		for _, num := range nums {
			inputChan <- num
		}
		close(inputChan)
	}()

	// pipe to double
	doubledChan := make(chan int, len(nums))
	go func() {
		defer close(doubledChan)
		for input := range inputChan {
			doubledChan <- double(input)
		}
	}()

	// pipe to addThree
	outputChan := make(chan int, len(nums))
	go func() {
		defer close(outputChan)
		for doubled := range doubledChan {
			outputChan <- addThree(doubled)
		}
	}()

	for output := range outputChan {
		fmt.Printf("Output: %d\n", output)
	}
}

func double(x int) int {
	return x * 2
}

func addThree(x int) int {
	return x + 3
}
