package main

import "fmt"

func main() {
	x := struct {
		first     string
		last      string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Mindaugas",
		last:  "Degutis",
		friends: map[string]int{
			"Vaida":     111,
			"Vidmantas": 222,
			"Lionytė":   333,
		},
		favDrinks: []string{
			"Coffee",
			"Wather",
		},
	}

	fmt.Println(x.first, x.last)
	fmt.Println(x.friends)
	fmt.Println(x.favDrinks)
}
