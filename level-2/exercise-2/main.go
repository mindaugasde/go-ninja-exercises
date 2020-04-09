package main

import "fmt"

func main() {
	a := 10
	b := 5

	de := (a == b * 2)
	le := (a <= b)
	ge := (a >= b)
	ne := (a != b)
	l := (a < b)
	g := (a > b)

	fmt.Println(de, le, ge, ne, l, g)
}
