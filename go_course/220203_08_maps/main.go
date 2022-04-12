// https://www.youtube.com/watch?v=rfQkR1bH3qw&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=76&ab_channel=AprendaGo

package main

import "fmt"

func main() {
	peopleHobby := map[string]string{
		"Jane_Mary":    "Jump",
		"Peter_Parker": "Walk",
		"Harry_Potter": "Clear",
	}

	for key, value := range peopleHobby {
		fmt.Println(key, value)
	}

	// Create a map where the keys are strings
	// and the values are slices of strings
	mapNew := map[string][]string{
		"bob_mcelreath": {"Jump", "Write", "Think about Bayesian Statistics"},
		"Brian_Kern":    {"Talk", "Do classes", "Write about C"},
		"Pensive_Cori":  {"Think about biology", "Be a nice container name"},
	}

	for key, value := range mapNew {
		fmt.Println(key, value)
	}
}

// next class
// https://www.youtube.com/watch?v=AKPdowl7tsw&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=77&ab_channel=AprendaGo
