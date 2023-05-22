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

	myTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: false,
	}

	mySedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: false,
	}

	fmt.Println("mySedan\t doors:", mySedan.doors, "\t color:", mySedan.color, "\t luxury:", mySedan.luxury)
	fmt.Println("myTruck\t doors:", myTruck.doors, "\t color:", myTruck.color, "\t fourWheel:", myTruck.fourWheel)

}
