package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Movie struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdb_id"`
	Type   string `json:"type"`
}

type Movies struct {
	Search       []Movie `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
}

func main() {
	http.HandleFunc("/movie", getMovieDetails)
	http.HandleFunc("/movies", getMovieList)
	http.HandleFunc("/movie/type", getMovieByTypeAndTitle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getMovieDetails(w http.ResponseWriter, r *http.Request) {
	imdbID := r.URL.Query().Get("imdbID")

	movie := Movie{
		Title:  "The True Story of Jesse James",
		Year:   "1957",
		ImdbID: imdbID,
		Type:   "movie",
	}

	// Convert the movie struct to JSON format
	jsonData, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}

func getMovieList(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	url := "http://www.omdbapi.com/?apikey=<your_api_key>&s=" + search + "&type=movie&page=" + page

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var movies Movies
	err = json.Unmarshal(body, &movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header
	w.Header().Set("Content-Type", "application/json")

	// Return the movie data as JSON
	jsonData, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func getMovieByTypeAndTitle(w http.ResponseWriter, r *http.Request) {
	movieType := r.URL.Query().Get("type")
	title := r.URL.Query().Get("search")

	// Here you can write code to retrieve the movie list from a database or an external API
	// For the sake of this example, we will return some static data

	url := "http://www.omdbapi.com/?apikey=<your_api_key>&s=" + title + "&type=" + strings.ToLower(movieType)

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var movies Movies
	err = json.Unmarshal(body, &movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header
	w.Header().Set("Content-Type", "application/json")

	// Return the movie data as JSON
	jsonData, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
