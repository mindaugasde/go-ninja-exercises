package main

import "fmt"

func main() {
	x := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", x)
}
