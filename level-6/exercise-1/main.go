package main

import "fmt"

func main() {
	b := foo()
	f, g := bar()
	fmt.Println(b, f, g)
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 15, "hell yes"
}
