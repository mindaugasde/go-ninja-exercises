package main

import "fmt"

func main() {
	stopper := func() int {
		return 1000000
	}

	fibonacciNumbers := fibonacci(stopper)
	fmt.Println(fibonacciNumbers)
}

func fibonacci(f func() int) []int {
	numbers := []int{0,1}
	lastIteration := f()
	for numbers[len(numbers)-1] < lastIteration {
		numbers = append(numbers, numbers[len(numbers)-1] + numbers[len(numbers)-2])
	}
	return numbers
}
