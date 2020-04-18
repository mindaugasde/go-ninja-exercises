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

	persons := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, person := range persons {
		fmt.Println(person.firstName)
		fmt.Println(person.lastName)
		for _, flavour := range person.iceCreamFlavors {
			fmt.Printf("\t%v\n", flavour)
		}
	}
}
