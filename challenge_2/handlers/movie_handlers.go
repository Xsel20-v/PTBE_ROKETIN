package handlers

import (
	"encoding/json"
	"net/http"
	"ptbe_roketin/challenge_2/models"
	"strconv"
	"strings"
)

var movies []models.Movie
var currentID = 1

//Create Movie
func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	movie.ID = currentID
	currentID++
	movies = append(movies, movie)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

// List all movies with pagination
func ListMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pageMovie := r.URL.Query().Get("page")
	limitMovie := r.URL.Query().Get("limit")

	// Default values
	page := 1
	limit := 10

	// Convert to integer
	if pageMovie != "" {
		p, err := strconv.Atoi(pageMovie)
		if err == nil && p > 0 {
			page = p
		}
	}

	if limitMovie != "" {
		l, err := strconv.Atoi(limitMovie)
		if err == nil && l > 0 {
			limit = l
		}
	}

	// set start and end
	start := (page - 1) * limit
	end := start + limit

	// if the start or end surpass the total item
	if start > len(movies) {
		start = len(movies)
	}
	if end > len(movies) {
		end = len(movies)
	}

	// pick within start & end
	paginatedMovies := movies[start:end]

	json.NewEncoder(w).Encode(paginatedMovies)
}

// Update movie
func UpdateMoviesHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	idMovie := r.URL.Query().Get("ID")
	id, err := strconv.Atoi((idMovie))
	if err != nil{
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
        return
	}

	var movieToUpdate *models.Movie
	for i := range movies {
		if movies[i].ID == id {
			movieToUpdate = &movies[i]
			break
		}
	}

	if movieToUpdate == nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
        return
	}

	var updatedMovie models.Movie
	err = json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	movieToUpdate.Title = updatedMovie.Title
	movieToUpdate.Description = updatedMovie.Description
	movieToUpdate.Artists = updatedMovie.Artists
	movieToUpdate.Duration = updatedMovie.Duration
	movieToUpdate.Genres = updatedMovie.Genres

	json.NewEncoder(w).Encode(movieToUpdate)
}

//Search movie by title/description/artist/genre
func SearchMoviesHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("q")
    if query == "" {
        http.Error(w, "Missing search query", http.StatusBadRequest)
        return
    }

    var results []models.Movie
    for _, movie := range movies {
        if strings.Contains(strings.ToLower(movie.Title), strings.ToLower(query)) ||
           strings.Contains(strings.ToLower(movie.Description), strings.ToLower(query)) ||
           containsIgnoreCase(movie.Artists, query) ||
           containsIgnoreCase(movie.Genres, query) {
            results = append(results, movie)
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}

// helper function for checking each item from Artist and Genre
func containsIgnoreCase(slice []string, str string) bool {
    str = strings.ToLower(str)
    for _, v := range slice {
        if strings.Contains(strings.ToLower(v), str) {
            return true
        }
    }
    return false
}
