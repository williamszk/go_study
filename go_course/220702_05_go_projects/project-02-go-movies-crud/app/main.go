package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func findMovieById(slice []Movie, id string) Movie {

	movie := Movie{}
	for i, v := range slice {
		if v.ID == id {
			movie = slice[i]
			break
		}
	}
	return movie
}

func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id := params["id"]

	movie := findMovieById(movies, id)

	json.NewEncoder(w).Encode(movie)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	log.Println(params)
	// params["id"]

	// for i, v := range movies {
	// 	if i.ID == params.id
	// }

	message := struct {
		Message string `json:"message"`
	}{
		Message: "Hello from delete",
	}
	// message: fmt.Sprintf("Message from the server ID: %s", string(params["id"])),

	json.NewEncoder(w).Encode(message)

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "8123712", Title: "Clean Code", Director: &Director{Firstname: "Bob", Lastname: "Martin"}})
	movies = append(movies, Movie{ID: "2", Isbn: "8123713", Title: "Unix Operating System", Director: &Director{Firstname: "Dennis", Lastname: "Ritchie"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
