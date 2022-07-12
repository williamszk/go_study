package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// build a command line tool to choose which request to make
	//

	func() {
		// ==================================== //
		// GET request for all movies
		// ==================================== //
		url := "http://localhost:8000/movies"

		res, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Response from get all movies")
		log.Println(string(body))
	}()

	func() {
		// ==================================== //
		// GET request for movie with id=1
		// ==================================== //
		url := "http://localhost:8000/movies/1"

		res, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Response from get movie with id=1")
		log.Println(string(body))
	}()

}
