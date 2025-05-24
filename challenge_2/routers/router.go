package routers

import (
	"ptbe_roketin/challenge_2/handlers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
    r := mux.NewRouter()

    r.HandleFunc("/movies", handlers.CreateMovieHandler).Methods("POST")
    r.HandleFunc("/movies", handlers.ListMoviesHandler).Methods("GET")
    r.HandleFunc("/movies", handlers.UpdateMoviesHandler).Methods("PUT")
    r.HandleFunc("/movies/search", handlers.SearchMoviesHandler).Methods("GET")

    return r
}