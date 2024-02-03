package themoviedb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TheMovieDBService struct {
	apiKey string
}

func NewTheMovieDBService(apiKey string) *TheMovieDBService {
	return &TheMovieDBService{
		apiKey: apiKey,
	}
}

func (s TheMovieDBService) GetMovieDetails(movieID string) (*MovieDetails, error) {

	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.themoviedb.org/3/movie/"+movieID+"?api_key="+s.apiKey, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var movieDetails MovieDetails
	if err := json.NewDecoder(res.Body).Decode(&movieDetails); err != nil {
		return nil, fmt.Errorf("GetMovieDetails: failed to decode response body: %w", err)
	}

	movieDetails.ImageURL = "https://image.tmdb.org/t/p/w500" + movieDetails.ImageURL

	return &movieDetails, nil
}

type MovieDetails struct {
	Title        string `json:"title"`
	Description  string `json:"overview"`
	IMDBID       string `json:"imdb_id"`
	TheMovieDBID int    `json:"id"`
	ImageURL     string `json:"poster_path"`
	// Need to add an unmarshaler for this since it's in 2014-10-22 format
	//ReleaseDate  time.Time `json:"release_date"`
}
