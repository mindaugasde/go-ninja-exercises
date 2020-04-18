package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName: "James",
		lastName:  "Bond",
		iceCreamFlavors: []string{
			"Chocolate",
			"Strawberry",
		},
	}

	p2 := person{
		firstName: "Miss",
		lastName:  "Moneypenny",
		iceCreamFlavors: []string{
			"Martini",
			"Cream",
		},
	}

	fmt.Println(p1)
	for _, flavour := range p1.iceCreamFlavors {
		fmt.Printf("\t%v\n", flavour)
	}
	fmt.Println(p2)
	for _, flavour := range p2.iceCreamFlavors {
		fmt.Printf("\t%v\n", flavour)
	}
}
