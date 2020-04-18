package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	truck1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println("Truck")
	fmt.Println("Doors:", truck1.doors)
	fmt.Println("Color:", truck1.color)
	fmt.Println("Four wheel:", truck1.fourWheel)

	fmt.Println("Sedan")
	fmt.Println("Doors:", sedan1.doors)
	fmt.Println("Color:", sedan1.color)
	fmt.Println("Luxury:", sedan1.luxury)
}
