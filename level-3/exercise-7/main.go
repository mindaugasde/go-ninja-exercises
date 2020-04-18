package main

import "fmt"

func main() {
	born := 1990
	if born <= 1960 {
		fmt.Println("You are in a very danger zone, please stay at home")
	} else if born <= 1990 {
		fmt.Println("You're good but please safe other people")
	} else {
		fmt.Println("Enjoy the life")
	}
}
