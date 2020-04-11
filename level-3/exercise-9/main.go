package main

import "fmt"

func main() {
	favSport := "auto"
	switch favSport {
	case "auto":
		fmt.Println("Favorite sport is auto sport")
	default:
		fmt.Println("None")
	}
}
