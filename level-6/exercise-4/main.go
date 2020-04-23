package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is,", p.first, p.last, "and I'm", p.age, "years young")
}

func main() {
	p1 := person{
		first: "Mindaugas",
		last:  "Degutis",
		age:   29,
	}

	p1.speak()
}
