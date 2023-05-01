package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", ISBN: "438227", Title: "Movie One", Director: &Director{Firstname: "Nana", Lastname: "Land"}})
	movies = append(movies, Movie{ID: "2", ISBN: "438228", Title: "Movie Two", Director: &Director{Firstname: "Lana", Lastname: "Land"}})

	router.HandleFunc("/api/movies", getMovies).Methods("GET")
	router.HandleFunc("/api/movies", createMovie).Methods("POST")
	router.HandleFunc("/api/movies", deleteMovies).Methods("DELETE")
	router.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server running in port at: 4000\n")
	log.Fatal(http.ListenAndServe(":4000", router))
}

func getMovies(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resw).Encode(movies)
}

func createMovie(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa(len(movies) + 1)
	movies = append(movies, movie)
	json.NewEncoder(resw).Encode(movie)
}

func deleteMovies(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-Type", "application/json")
	movies = []Movie{}
	json.NewEncoder(resw).Encode(movies)
}

func getMovie(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(resw).Encode(item)
			return
		}
	}
	json.NewEncoder(resw).Encode("No movie found!")
}

func updateMovie(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, item := range movies {
		if item.ID == params["id"] {
			// movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)
			item.Title = movie.Title
			item.ISBN = movie.ISBN
			item.Director.Firstname = movie.Director.Firstname
			item.Director.Lastname = movie.Director.Lastname
			json.NewEncoder(resw).Encode(item)
			break
		}
	}
}

func deleteMovie(resw http.ResponseWriter, req *http.Request) {
	resw.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(resw).Encode(movies)
}
