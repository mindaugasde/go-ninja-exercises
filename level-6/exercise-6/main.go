package main

import "fmt"

func main() {
	x := func() string {
		return "Hello Ladies"
	}()

	fmt.Println(x)
}
