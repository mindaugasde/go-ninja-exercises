package main

import "fmt"

func main() {
	f := func(fn string) {
		fmt.Println("My name is", fn)
	}

	f("Mindaugas")
}
