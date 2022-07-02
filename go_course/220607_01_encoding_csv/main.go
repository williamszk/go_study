// https://www.youtube.com/watch?v=VQnbTyzeNic&ab_channel=code_with_kavit
// Objective is to save files to csv
// we use os to create a csv file

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

	f, err := os.Create("users.csv")
	// defer will be executed in the end of the program
	// that is, when the function main ends
	defer f.Close()

	if err != nil {
		log.Fatalln("Failed to open file: ", err)
		// fmt.Println(err)
	}

	// create a writter object
	w := csv.NewWriter(f)
	// the new writer will add data to the buffer
	// we need to use flush so that the data inside the
	// buffer will go to the file
	// it take less time to write to a buffer
	// and then flush it into the file
	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("Error writing record to file: ", err)
		}
	}
}
