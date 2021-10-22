package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shivamtiwari/database"
	"github.com/shivamtiwari/models"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var movie models.Movie
	err := decoder.Decode(&movie)
	if err != nil {
		log.Println(err.Error())
	}
	database.MovieList = append(database.MovieList, movie)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	allMovie := database.MovieList
	json.NewEncoder(w).Encode(allMovie)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id := mux.Vars(r)["id"]

	for _, movie := range database.MovieList {
		if id == movie.Id {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	message := models.Message{Message: "Movie Not Found"}
	json.NewEncoder(w).Encode(message)
}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id := mux.Vars(r)["id"]

	for index, movie := range database.MovieList {
		if id == movie.Id {
			database.MovieList = append(database.MovieList[:index], database.MovieList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(database.MovieList)
}

func UpdateMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)

	for index, movieToUpdate := range database.MovieList {
		if movieToUpdate.Id == movie.Id {
			database.MovieList[index] = movie
		}
	}
}
