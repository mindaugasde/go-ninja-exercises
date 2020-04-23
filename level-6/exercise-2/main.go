package main

import "fmt"

func main() {
	shots := []int{2, 4, 7, 9}
	sum := foo(shots...)
	fmt.Println(sum)

	sum2 := bar(shots)
	fmt.Println(sum2)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
