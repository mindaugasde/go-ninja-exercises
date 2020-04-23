package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	defer func() {
		fmt.Println("exit...")
	}()
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
