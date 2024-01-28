package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	moviesv1 "github.com/griggsca91/givemeadamnmovie/gen/go/movies/v1"
	"github.com/griggsca91/givemeadamnmovie/gen/go/movies/v1/moviesv1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}

type MovieServer struct {
}

// GetMovieRecommendation implements moviesv1connect.MoviesServiceHandler.
func (*MovieServer) GetMovieRecommendation(context.Context, *connect.Request[moviesv1.GetMovieRecommendationRequest]) (*connect.Response[moviesv1.GetMovieRecommendationResponse], error) {
	fmt.Println("Hit")
	return &connect.Response[moviesv1.GetMovieRecommendationResponse]{
		Msg: &moviesv1.GetMovieRecommendationResponse{
			Movie: &moviesv1.Movie{
				Description: "farts",
				Title:       "farts",
			},
		},
	}, nil
}

func main() {
	movieServer := MovieServer{}
	mux := http.NewServeMux()

	path, handler := moviesv1connect.NewMoviesServiceHandler(&movieServer)

	mux.Handle(path, withCORS(handler))
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

}
