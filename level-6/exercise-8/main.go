package main

import "fmt"

func main() {
	f := factorial(5)
	fmt.Println(f())
}

func factorial(n int) func() int {
	return func() int {
		fact := 1
		for ; n > 0; n-- {
			fact *= n
		}
		return fact
	}
}
