package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	xp := [][]string{jb, mp}
	for i, v := range xp {
		fmt.Println(i, v)
		for j, vv := range v {
			fmt.Println(j, vv)
		}
	}
}
