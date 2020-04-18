package main

import "fmt"

func main() {

	const (
		currentYear = 2020
		a           = currentYear + iota
		b           = currentYear + iota
		c           = currentYear + iota
		d           = currentYear + iota
	)

	fmt.Println(a, b, c, d)
}
