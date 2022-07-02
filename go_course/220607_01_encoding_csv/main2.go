// here we use another methods writeall to write everything of an object
// inside the csv, if the data is already in a 2-d format

package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "occupation"},
		{"John", "Doe", "Paiter"},
		{"Lucy", "Smith", "Programmer"},
		{"Nick", "Fury", "Director of SHIELD"},
	}

	f, err := os.Create("users1.csv")
	defer f.Close()

	if err != nil {
		log.Fatalln("Failed to open file: ", err)
	}

	w := csv.NewWriter(f)

	// the method WriteAll already have a flush
	err = w.WriteAll(records)
	if err != nil {
		log.Fatalln(err)
	}

	// defer w.Flush()

	// for _, record := range records {
	// 	if err := w.Write(record); err != nil {
	// 		log.Fatalln("Error writing record to file: ", err)
	// 	}
	// }
}
