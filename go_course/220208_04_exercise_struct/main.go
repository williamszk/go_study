package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	premiumModel bool
}

func main() {

	veh01 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "green",
		},
		fourWheels: false,
	}

	veh02 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		premiumModel: true,
	}

	fmt.Println("Color of vehicle 01:", veh01.vehicle.color)
	fmt.Println("Is vehicle 02 premium model?:", veh02.premiumModel)

}
