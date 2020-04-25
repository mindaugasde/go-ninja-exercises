package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person, fn string, ln string) {
	p.first = fn
	p.last = ln
}

func main() {
	p1 := person{
		first: "Mindaugas",
		last:  "Degutis",
	}
	fmt.Println(p1)

	changeMe(&p1, "John", "Doe")
	fmt.Println(p1)
}
