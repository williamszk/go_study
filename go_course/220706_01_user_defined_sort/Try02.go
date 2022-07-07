package main

import "fmt"

// Name: name of the airplane model
// Used: Is a Union of string it can have "Military" or "Civilian"
// Year: year of development
// Range: Maximum range in kilometers
type Airplane struct {
	Name  string
	Used  string
	Year  int
	Range float64
}

type AirplaneList []Airplane

func (a AirplaneList) Len() int {
	return len(a)
}

// func (a AirplaneList) Less(i, j int) bool {

// }

// func (a AirplaneList) Swap(i, j int) {

// }

func Try02() {

	airplaneList := AirplaneList(
		[]Airplane{
			{
				Name:  "F-15",
				Used:  "Military",
				Year:  1976,
				Range: 4815.0,
			},
			{
				Name:  "Boeing 787",
				Used:  "Civilian",
				Year:  2011,
				Range: 13530.0,
			},
			{
				Name:  "Curtiss P-40 Warhawk",
				Used:  "Military",
				Year:  1938,
				Range: 386.0,
			},
			{
				Name:  "Lockheed SR-71 Blackbird",
				Used:  "Military",
				Year:  1966,
				Range: 5400.0,
			},
		},
	)

	fmt.Println(airplaneList)

	fmt.Println(airplaneList.Len())

	// Sort the airplanes by year of development

}
