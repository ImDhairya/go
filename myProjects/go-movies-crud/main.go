package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updatedMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
	params := mux.Vars(r)
	id := params["id"]

	// movieId := movie.ID

	for index, item := range movies {
		if item.ID == params["id"] {
			movies[index] = updatedMovie
			// but make sure the ID doesn’t change
			movies[index].ID = id

			json.NewEncoder(w).Encode(movies[index])
			return

		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func main() {
	fmt.Println("Hii")

	movies = append(movies, Movie{"1", "438227", "The Redemption", &Director{Fname: "Dhairya", Lname: "Pandya"}})

	movies = append(movies, Movie{"2", "432321", "The Shining", &Director{"Chinmay", "Nagar"}})

	movies = append(movies, Movie{"3", "5348898", "Fast and Furious", &Director{"Abhishek", "Pandya"}})

	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")

	log.Fatal(http.ListenAndServe(":8000", r))

}
