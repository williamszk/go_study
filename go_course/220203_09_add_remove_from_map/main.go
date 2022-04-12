// https://www.youtube.com/watch?v=AKPdowl7tsw&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=78&ab_channel=AprendaGo
package main

import "fmt"

func main() {
	mapNew := map[string][]string{
		"bob_mcelreath": {"Jump", "Write", "Think about Bayesian Statistics"},
		"Brian_Kern":    {"Talk", "Do classes", "Write about C"},
		"Pensive_Cori":  {"Think about biology", "Be a nice container name"},
	}

	// add new item into the map
	mapNew["knowledgeable_lucas"] = []string{"Skywalker", "Obi-wan", "I'm your father"}

	for key, value := range mapNew {
		fmt.Println(key, value)
	}

	fmt.Println("")

	// remove item from map
	delete(mapNew, "Brian_Kern")

	for key, value := range mapNew {
		fmt.Println(key, value)
	}

}

// next:
// https://www.youtube.com/watch?v=EaOGcmXo4F8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=79&ab_channel=AprendaGo
