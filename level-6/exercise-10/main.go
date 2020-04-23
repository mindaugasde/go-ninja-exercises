package main

import "fmt"

func main() {
	bob := workTime()
	john := workTime()

	fmt.Println(bob(1.0))
	fmt.Println(bob(2.5))
	fmt.Println(bob(12.3))
	fmt.Println(john(0.5))
	fmt.Println(john(7.1))
}

func workTime() func(wt float64) float64 {
	x := 0.0
	return func(wt float64) float64 {
		x += wt
		return x
	}
}
